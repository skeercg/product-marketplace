package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"product-marketplace/pkg/model"
	tokenService "product-marketplace/pkg/service"
)

type ProductController struct {
	tokenService.ProductService
}

func (controller *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params model.SearchProductParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Print(err)
	}

	params.Token = r.Header.Get("Authorization")

	if tokenService.IsTokenExpired(params.Token) {
		w.WriteHeader(401)
		return
	}

	products := controller.ProductService.SearchProduct(params)

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Print(err)
	}
}

func (controller *ProductController) GradeProduct(w http.ResponseWriter, r *http.Request) {
	var params model.GradeProductParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Print(err)
		w.WriteHeader(400)
	}

	params.Token = r.Header.Get("Authorization")

	if tokenService.IsTokenExpired(params.Token) {
		w.WriteHeader(401)
		return
	}

	err = controller.ProductService.GradeProduct(params)

	if err != nil {
		w.WriteHeader(500)
	}
}
