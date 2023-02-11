package _interface

import (
	en "product-marketplace/pkg/model"
)

type ProductRepository interface {
	SearchProduct(ProductPriceFrom, ProductPriceTo int, ProductName string, ProductRating float32) []en.Product

	GradeProduct(ProductId int, ProductRating float32) error
}
