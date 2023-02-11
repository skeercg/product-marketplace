package _interface

import "product-marketplace/pkg/service/interface"

type AuthRepository interface {
	_interface.TokenService

	AuthorizeUser(UserName, Password string) (string, error)

	CreateUser(UserName, Password string) (string, error)
}
