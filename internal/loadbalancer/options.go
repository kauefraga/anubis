package loadbalancer

import "github.com/kauefraga/anubis/internal/algorithms"

func WithPort(port uint16) LoadBalancerOption {
	return func(lb *LoadBalancer) {
		lb.Port = port
	}
}

func WithServers(servers []string) LoadBalancerOption {
	return func(lb *LoadBalancer) {
		lb.Servers = servers
	}
}

func WithAlgorithm(algorithm string) LoadBalancerOption {
	return func(lb *LoadBalancer) {
		lb.Algorithm = algorithms.RoundRobin()

		// TODO add least connection and weighted response time algorithms
	}
}
