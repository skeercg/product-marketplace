package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-marketplace/pkg/init"
)

func main() {
	init.InitDbConnection()
	router := mux.NewRouter().StrictSlash(true)

	init.InitProductApi(router)
	init.InitAuthApi(router)

	err := http.ListenAndServe(":8181", router)
	if err != nil {
		log.Print(err)
	}
}
