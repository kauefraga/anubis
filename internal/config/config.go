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

func Read() Config {
	doc, err := os.ReadFile("anubis.toml")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	var cfg Config

	err = toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	if cfg.Version != 1 {
		cfg.Version = 1
	}

	if cfg.Port == 0 {
		cfg.Port = 4000
	}

	if cfg.Algorithm != "round-robin" && cfg.Algorithm != "least-connection" && cfg.Algorithm != "weighted-response-time" {
		cfg.Algorithm = "round-robin"
	}

	if len(cfg.Servers) == 0 {
		log.Fatalln("Error: there are no servers in the anubis.toml")
	}

	return cfg
}
