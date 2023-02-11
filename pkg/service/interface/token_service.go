package _interface

type TokenService interface {
	GenerateToken() (string, error)

	IsTokenExpired(string) bool
}
