package config

import (
	"log"
	"os"

	"github.com/kauefraga/anubis/internal/models"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Version   int
	Port      uint16
	Algorithm string
	Servers   []*models.Servers
}

func Read() *Config {
	doc, err := os.ReadFile("anubis.toml")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	var cfg Config

	err = toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	return &cfg
}
