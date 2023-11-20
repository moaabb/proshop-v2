package config

import (
	"log"
	"os"
)

type Config struct {
	Port             string
	AuthSvcUrl       string
	SaveFileBasePath string
}

func NewConfig() (c *Config) {
	c = &Config{
		Port:             os.Getenv("PORT"),
		AuthSvcUrl:       os.Getenv("AUTH_SVC_URL"),
		SaveFileBasePath: os.Getenv("SAVE_FILE_BASE_PATH"),
	}

	if c.Port == "" || c.AuthSvcUrl == "" || c.SaveFileBasePath == "" {
		log.Fatal("could not get env")
	}
	return
}
