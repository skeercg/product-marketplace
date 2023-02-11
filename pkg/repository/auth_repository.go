package repository

type AuthRepository interface {
	AuthorizeUser(UserName, Password string) (string, error)

	CreateUser(UserName, Password string) (string, error)
}
