package setup

import (
	"github.com/gorilla/mux"
	"product-marketplace/pkg/controller"
	implementation2 "product-marketplace/pkg/repository/implementation"
	"product-marketplace/pkg/service/implementation"
	sr "product-marketplace/pkg/service/implementation"
)

func AuthControllerSetup(router *mux.Router) {
	tokenService := new(sr.TokenServiceImpl)
	repository := new(implementation2.AuthRepositoryImpl)
	repository.TokenService = tokenService
	service := new(implementation.AuthServiceImpl)
	service.AuthRepository = repository
	controller := new(controller.AuthController)
	controller.AuthService = service

	s := router.PathPrefix("/accounts").Subrouter()

	s.HandleFunc("/login", controller.LoginUser).Methods("POST")

	s.HandleFunc("/register", controller.RegisterUser).Methods("POST")
}
