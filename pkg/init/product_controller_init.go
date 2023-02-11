package init

import (
	"github.com/gorilla/mux"
	storePresentation "product-marketplace/pkg/controller"
	storeData "product-marketplace/pkg/repository"
	storeDomain "product-marketplace/pkg/service"
)

func InitProductApi(router *mux.Router) {
	repository := new(storeData.ProductRepositoryImpl)
	service := new(storeDomain.ProductServiceImpl)
	service.ProductRepository = repository
	controller := new(storePresentation.ProductController)
	controller.ProductService = service

	s := router.PathPrefix("/products").Subrouter()

	s.HandleFunc("/search", controller.GetProducts).Methods("POST")

	s.HandleFunc("/grade", controller.GradeProduct).Methods("POST")
}
