package implementation

import (
	"product-marketplace/pkg/model"
	"product-marketplace/pkg/repository/interface"
)

type ProductServiceImpl struct {
	_interface.ProductRepository
}

func (p *ProductServiceImpl) SearchProduct(Params model.SearchProductParams) []model.Product {
	return p.ProductRepository.SearchProduct(Params.ProductPriceFrom, Params.ProductPriceTo, Params.ProductName, Params.ProductRatingFrom)
}

func (p *ProductServiceImpl) GradeProduct(Params model.GradeProductParams) error {
	return p.ProductRepository.GradeProduct(Params.ProductId, Params.Rating)
}
