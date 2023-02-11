package service

import (
	"product-marketplace/pkg/model"
	"product-marketplace/pkg/repository"
)

type ProductServiceImpl struct {
	repository.ProductRepository
}

func (p *ProductServiceImpl) SearchProduct(Params model.SearchProductParams) []model.Product {
	return p.ProductRepository.SearchProduct(Params.ProductPriceFrom, Params.ProductPriceTo, Params.ProductName, Params.ProductRatingFrom)
}

func (p *ProductServiceImpl) GradeProduct(Params model.GradeProductParams) error {
	return p.ProductRepository.GradeProduct(Params.ProductId, Params.Rating)
}
