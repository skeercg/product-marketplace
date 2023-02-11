package setup

import (
	"github.com/gorilla/mux"
	storePresentation "product-marketplace/pkg/controller"
	storeData "product-marketplace/pkg/repository/implementation"
	storeDomain "product-marketplace/pkg/service/implementation"
)

func ProductControllerSetup(router *mux.Router) {
	tokenService := new(storeDomain.TokenServiceImpl)
	repository := new(storeData.ProductRepositoryImpl)
	service := new(storeDomain.ProductServiceImpl)
	service.ProductRepository = repository
	controller := new(storePresentation.ProductController)
	controller.ProductService = service
	controller.TokenService = tokenService

	s := router.PathPrefix("/products").Subrouter()

	s.HandleFunc("/search", controller.GetProducts).Methods("POST")

	s.HandleFunc("/grade", controller.GradeProduct).Methods("POST")
}
