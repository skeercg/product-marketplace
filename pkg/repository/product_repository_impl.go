package repository

import (
	"fmt"
	"log"
	"product-marketplace/pkg/init"
	en "product-marketplace/pkg/model"
)

type ProductRepositoryImpl struct {
}

func (p *ProductRepositoryImpl) SearchProduct(ProductPriceFrom, ProductPriceTo int, ProductName string, ProductRatingFrom float32) []en.Product {
	query := fmt.Sprintf(`select p.id, p.name, p.price, (pr.total_rating / pr.rating_count) as rating 
								from products as p, products_rating as pr
								where p.id = pr.id and $1 <= p.price and p.price <= $2 
								  and $3 <= (pr.total_rating / pr.rating_count) and lower(p.name) like $4`)

	nameRegex := "%" + fmt.Sprintf("%s", ProductName) + "%"

	rows, err := init.DBInstance.Query(query, ProductPriceFrom, ProductPriceTo, ProductRatingFrom, nameRegex)
	if err != nil {
		log.Print(err)
	}

	var products []en.Product

	for rows.Next() {
		p := en.Product{}

		err = rows.Scan(&p.ProductId, &p.ProductName, &p.ProductPrice, &p.ProductRating)

		if err != nil {
			log.Print(err)
		}

		products = append(products, p)
	}

	return products
}

func (p *ProductRepositoryImpl) GradeProduct(ProductId int, ProductRating float32) error {
	query := fmt.Sprintf(`update products_rating 
						set total_rating = total_rating + $1, rating_count = rating_count + 1 
						where id = $2`)

	_, err := init.DBInstance.Exec(query, ProductRating, ProductId)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
