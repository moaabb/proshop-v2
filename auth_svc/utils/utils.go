package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/moaabb/ecommerce/auth_svc/domain/user"
)

type CustomClaims struct {
	*jwt.RegisteredClaims
	IsAdmin bool
	UserId  uint
}

func GenerateJWT(u user.User, secret string) (string, CustomClaims, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "proshop",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		IsAdmin: u.IsAdmin,
		UserId:  u.Id,
	})
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", CustomClaims{}, err
	}

	return ss, token.Claims.(CustomClaims), nil

}

func VerifyJWT(token string, secret string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid Token")

}
