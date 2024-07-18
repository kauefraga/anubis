package loadbalancer

import (
	"github.com/kauefraga/anubis/internal/algorithms"
	"github.com/kauefraga/anubis/internal/models"
)

func WithPort(port uint16) LoadBalancerOption {
	return func(lb *LoadBalancer) {
		lb.Port = port
	}
}

func WithServers(servers []*models.Server) LoadBalancerOption {
	return func(lb *LoadBalancer) {
		lb.Servers = servers
	}
}

func WithAlgorithm(algorithm string) LoadBalancerOption {
	return func(lb *LoadBalancer) {
		if algorithm == "round-robin" {
			lb.Algorithm = algorithms.RoundRobin()
		}

		if algorithm == "least-connection" {
			lb.Algorithm = algorithms.LeastConnection()
		}
	}
}
