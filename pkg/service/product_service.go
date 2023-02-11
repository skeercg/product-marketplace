package service

import (
	"product-marketplace/pkg/model"
)

type ProductService interface {
	SearchProduct(model.SearchProductParams) []model.Product

	GradeProduct(model.GradeProductParams) error
}
