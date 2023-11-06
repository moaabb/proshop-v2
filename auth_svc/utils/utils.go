package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId uint, secret string) (string, jwt.RegisteredClaims, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "proshop",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		ID:        fmt.Sprint(userId),
	})
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", jwt.RegisteredClaims{}, err
	}

	return ss, token.Claims.(jwt.RegisteredClaims), nil

}

func VerifyJWT(token string, secret string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := t.Claims.(*jwt.RegisteredClaims); ok && t.Valid {
		return claims.ID, nil
	}

	return "", fmt.Errorf("invalid Token")

}
