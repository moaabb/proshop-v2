package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
	Dsn  string
}

func NewConfig() (c *Config) {
	c = &Config{
		Port: os.Getenv("PORT"),
		Dsn:  os.Getenv("DB_URL"),
	}

	if c.Dsn == "" || c.Port == "" {
		log.Fatal("could not get env")
	}
	return
}
