package config

import (
	"log"
	"os"
)

type Config struct {
	Port       string
	Dsn        string
	AuthSvcUrl string
}

func NewConfig() (c *Config) {
	c = &Config{
		Port:       os.Getenv("PORT"),
		Dsn:        os.Getenv("DB_URL"),
		AuthSvcUrl: os.Getenv("AUTH_SVC_URL"),
	}

	if c.Dsn == "" || c.Port == "" || c.AuthSvcUrl == "" {
		log.Fatal("could not get env")
	}
	return
}
