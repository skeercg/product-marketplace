package _interface

// AuthService Service that resides authorization
// business logic
type AuthService interface {
	LoginUser(UserName, Password string) (string, error)

	RegisterUser(UserName, Password string) (string, error)
}
