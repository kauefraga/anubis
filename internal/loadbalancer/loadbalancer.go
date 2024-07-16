package loadbalancer

import (
	"fmt"
	"net/http"

	"github.com/kauefraga/anubis/internal/algorithms"
)

type LoadBalancer struct {
	Algorithm algorithms.Algorithm
	Servers   []string
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
		ch := make(chan string)

		go func(ch chan string) {
			server := lb.Algorithm(lb.Servers)

			fmt.Println("Proxy request to", server) // TODO: remove this debug

			ch <- server
		}(ch)

		server, ok := <-ch

		if ok {
			close(ch)
			http.Redirect(w, r, server, http.StatusSeeOther)
		}
	})

	fmt.Printf("Listening on http://localhost:%d\n", lb.Port)
	return http.ListenAndServe(fmt.Sprint(":", lb.Port), nil)
}
