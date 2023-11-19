package config

import (
	"log"
	"os"
)

type Config struct {
	Port           string
	Dsn            string
	JwtSecret      string
	PaypalClientId string
	Domain         string
	SecureCookie   bool
}

func NewConfig() (c *Config) {
	var sc = false

	if os.Getenv("SECURE_COOKIE") == "true" {
		sc = true
	}

	c = &Config{
		Port:           os.Getenv("PORT"),
		Dsn:            os.Getenv("DB_URL"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		PaypalClientId: os.Getenv("PAYPAL_CLIENT_ID"),
		Domain:         os.Getenv("DOMAIN"),
		SecureCookie:   sc,
	}

	if c.Dsn == "" || c.JwtSecret == "" || c.Port == "" || c.PaypalClientId == "" || c.Domain == "" {
		log.Fatal("could not get env")
	}
	return
}
