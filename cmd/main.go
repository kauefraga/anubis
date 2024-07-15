package main

import (
	"log"

	"github.com/kauefraga/anubis/internal/config"
	"github.com/kauefraga/anubis/internal/loadbalancer"
)

func main() {
	cfg := config.Read()

	lb := loadbalancer.New(
		loadbalancer.WithPort(cfg.Port),
		loadbalancer.WithServers(cfg.Servers),
		loadbalancer.WithAlgorithm(cfg.Algorithm),
	)

	log.Fatalln(lb.Listen())
}
