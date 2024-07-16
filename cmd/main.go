package main

import (
	"fmt"
	"log"

	"github.com/kauefraga/anubis/internal/config"
	"github.com/kauefraga/anubis/internal/loadbalancer"
)

func main() {
	cfg := config.Read()

	if len(cfg.Servers) == 1 {
		fmt.Println("Warn: there is no advantage in using a load balancer with just one server")
	}

	lb := loadbalancer.New(
		loadbalancer.WithPort(cfg.Port),
		loadbalancer.WithServers(cfg.Servers),
		loadbalancer.WithAlgorithm(cfg.Algorithm),
	)

	log.Fatalln(lb.Listen())
}
