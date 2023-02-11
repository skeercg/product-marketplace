package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"product-marketplace/pkg/model"
	sr "product-marketplace/pkg/service"
)

type AuthController struct {
	sr.AuthService
}

func (controller *AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u model.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Print(err)
	}

	token, err := controller.AuthService.LoginUser(u.Username, u.Password)

	if err != nil {
		w.WriteHeader(401)
		return
	}

	json.NewEncoder(w).Encode(token)
}

func (controller *AuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u model.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Print(err)
	}

	token, err := controller.AuthService.RegisterUser(u.Username, u.Password)

	if err != nil {
		w.WriteHeader(401)
		return
	}

	json.NewEncoder(w).Encode(token)
}
