package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Version   int
	Port      uint16
	Algorithm string
	Servers   []string
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
		log.Fatalln("Error: this Anubis version does not exist")
	}

	if len(cfg.Servers) == 0 {
		log.Fatalln("Error: there are no servers in the anubis.toml")
	}

	if cfg.Algorithm != "round-robin" && cfg.Algorithm != "least-connection" && cfg.Algorithm != "weighted-response-time" {
		log.Fatalln("Error: this algorithm is not implemented")
	}

	return cfg
}
