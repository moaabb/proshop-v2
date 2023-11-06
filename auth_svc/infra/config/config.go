package config

import (
	"log"
	"os"
)

type Config struct {
	Port      string
	Dsn       string
	JwtSecret string
}

func NewConfig() (c *Config) {
	c = &Config{
		Port:      os.Getenv("PORT"),
		Dsn:       os.Getenv("DB_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	if c.Dsn == "" || c.JwtSecret == "" || c.Port == "" {
		log.Fatal("could not get env")
	}
	return
}
