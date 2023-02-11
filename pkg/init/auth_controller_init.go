package init

import (
	"github.com/gorilla/mux"
	"product-marketplace/pkg/controller"
	"product-marketplace/pkg/repository"
	"product-marketplace/pkg/service"
)

func InitAuthApi(router *mux.Router) {
	repository := new(repository.AuthRepositoryImpl)
	service := new(service.AuthServiceImpl)
	service.AuthRepository = repository
	controller := new(controller.AuthController)
	controller.AuthService = service

	s := router.PathPrefix("/accounts").Subrouter()

	s.HandleFunc("/login", controller.LoginUser).Methods("POST")

	s.HandleFunc("/register", controller.RegisterUser).Methods("POST")
}
