package config

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func isServerValid(server string) bool {
	s := strings.ToLower(server)

	if len(s) == 0 {
		return false
	}

	url, err := url.Parse(s)
	if err != nil || url.Host == "" {
		return false
	}

	if url.Scheme == "" && url.Fragment == "" && url.Opaque == "" {
		return false
	}

	return url.Scheme == "http" || url.Scheme == "https"
}

func (cfg *Config) Validate() *Config {
	if len(cfg.Servers) == 0 {
		log.Fatalln("Error: there are no servers in the anubis.toml")
	}

	if len(cfg.Servers) == 1 {
		fmt.Println("Warn: there is no advantage in using a load balancer with just one server")
	}

	for _, s := range cfg.Servers {
		if !isServerValid(s) {
			log.Fatalln("Error: the server", s, "is invalid")
		}
	}

	return cfg
}
