package config

import (
	"log"
	"os"
)

type Config struct {
	Port           string
	Dsn            string
	PaypalClientId string
	PaypalSecretId string
	PaypalBaseUrl  string
}

func NewConfig() (c *Config) {
	c = &Config{
		Port:           os.Getenv("PORT"),
		Dsn:            os.Getenv("DB_URL"),
		PaypalClientId: os.Getenv("PAYPAL_CLIENT_ID"),
		PaypalSecretId: os.Getenv("PAYPAL_SECRET_ID"),
		PaypalBaseUrl:  os.Getenv("PAYPAL_BASE_URL"),
	}

	if c.Dsn == "" || c.Port == "" || c.PaypalBaseUrl == "" || c.PaypalClientId == "" || c.PaypalSecretId == "" {
		log.Fatal("could not get env")
	}
	return
}
