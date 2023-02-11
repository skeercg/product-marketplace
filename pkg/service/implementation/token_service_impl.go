package implementation

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"log"
	"strings"
	"time"
)

var JwtKey = []byte("23c9b16565955d2461a1103fbbdbffb9")

type TokenServiceImpl struct {
}

func (service *TokenServiceImpl) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	})

	return token.SignedString(JwtKey)
}

func (service *TokenServiceImpl) IsTokenExpired(header string) bool {
	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		log.Print("invalid token length")
		return true
	}

	token, err := jwt.ParseWithClaims(headerParts[1], &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return JwtKey, nil
	})

	if err != nil {
		log.Print(err)
		return true
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)

	if !ok {
		log.Print("invalid claims")
		return true
	}

	return claims.ExpiresAt < time.Now().Unix()
}
