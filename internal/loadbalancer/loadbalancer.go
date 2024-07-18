package loadbalancer

import (
	"fmt"
	"net/http"

	"github.com/kauefraga/anubis/internal/algorithms"
	"github.com/kauefraga/anubis/internal/models"
)

type LoadBalancer struct {
	Algorithm algorithms.Algorithm
	Servers   []*models.Servers
	Port      uint16
}

type LoadBalancerOption func(*LoadBalancer)

func New(opts ...LoadBalancerOption) *LoadBalancer {
	lb := &LoadBalancer{
		Algorithm: algorithms.RoundRobin(),
		Servers:   nil,
		Port:      4000,
	}

	for _, setOpt := range opts {
		setOpt(lb)
	}

	return lb
}

func (lb *LoadBalancer) Listen() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lbServers := lb.Servers
		s := lb.Algorithm(lbServers)

		// You will need parse the url string to url.Url type
		// Then use httputtil.ReverseProxy(url).Serve(w,r)

		// I thought about creating a temp config, parse the string and then
		// pass the value for the final config
		fmt.Fprintln(w, "You are on the server", s)
	})

	fmt.Printf("Listening on http://localhost:%d\n", lb.Port)
	return http.ListenAndServe(fmt.Sprint(":", lb.Port), nil)
}
