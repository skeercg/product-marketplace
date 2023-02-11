package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	init2 "product-marketplace/pkg/controller/setup"
	"product-marketplace/pkg/repository/setup"
)

func main() {
	setup.DBConnectionSetup()
	router := mux.NewRouter().StrictSlash(true)

	init2.ProductApiSetup(router)
	init2.AuthApiSetup(router)

	err := http.ListenAndServe(":8181", router)
	if err != nil {
		log.Print(err)
	}
}
