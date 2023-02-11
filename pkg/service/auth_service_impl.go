package service

import (
	"log"
	"product-marketplace/pkg/repository"
)

type AuthServiceImpl struct {
	repository.AuthRepository
}

func (a *AuthServiceImpl) LoginUser(UserName, Password string) (string, error) {
	token, err := a.AuthRepository.AuthorizeUser(UserName, Password)

	if err != nil {
		log.Print(err)
		return "", err
	}

	return token, nil
}

func (a *AuthServiceImpl) RegisterUser(UserName, Password string) (string, error) {
	token, err := a.AuthRepository.CreateUser(UserName, Password)

	if err != nil {
		log.Print(err)
		return "", err
	}

	return token, nil
}
