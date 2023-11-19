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
}

func NewConfig() (c *Config) {
	c = &Config{
		Port:           os.Getenv("PORT"),
		Dsn:            os.Getenv("DB_URL"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		PaypalClientId: os.Getenv("PAYPAL_CLIENT_ID"),
	}

	if c.Dsn == "" || c.JwtSecret == "" || c.Port == "" || c.PaypalClientId == "" {
		log.Fatal("could not get env")
	}
	return
}
