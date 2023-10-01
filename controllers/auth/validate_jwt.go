package authcontroller

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ValidateJwt(tokenStr string) (*CustomJwtClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomJwtClaims{}, keyFunc)

	claims, ok := token.Claims.(*CustomJwtClaims)

	if err != nil {
		return nil, err
	} else if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, err
}
