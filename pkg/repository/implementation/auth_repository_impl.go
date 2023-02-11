package implementation

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"product-marketplace/pkg/model"
	"product-marketplace/pkg/repository/setup"
	"product-marketplace/pkg/service"
)

type AuthRepositoryImpl struct{}

func (a *AuthRepositoryImpl) AuthorizeUser(UserName, Password string) (string, error) {
	query := fmt.Sprintf("select * from users where username = $1 and password = $2")

	var u model.User
	err := setup.DBInstance.Get(&u, query, UserName, Password)

	if err != nil {
		log.Print(err)
		return "", err
	}

	return service.GenerateToken()
}

func (a *AuthRepositoryImpl) CreateUser(UserName, Password string) (string, error) {
	query := fmt.Sprintf("insert into users (username, password) values ($1, $2) returning *")

	row := setup.DBInstance.QueryRowx(query, UserName, Password)

	var u, p interface{}

	if err := row.Scan(&u, &p); err != nil {
		log.Print(err)
		return "", err
	}

	return service.GenerateToken()
}
