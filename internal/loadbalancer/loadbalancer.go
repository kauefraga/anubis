package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kauefraga/anubis/internal/algorithms"
	"github.com/kauefraga/anubis/internal/models"
)

type LoadBalancer struct {
	Algorithm algorithms.Algorithm
	Servers   []*models.Server
	Port      uint16
}

type LoadBalancerOption func(*LoadBalancer)

func New(opts ...LoadBalancerOption) *LoadBalancer {
	lb := &LoadBalancer{
		Algorithm: nil,
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

		u, err := url.Parse(s.Url)
		if err != nil {
			log.Println("Error: failed parsing server URL")
		}

		httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
	})

	fmt.Printf("Listening on http://localhost:%d\n", lb.Port)
	return http.ListenAndServe(fmt.Sprint(":", lb.Port), nil)
}
