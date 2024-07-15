package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kauefraga/anubis/internal/algorithms"
	"github.com/kauefraga/anubis/internal/config"
)

func main() {
	cfg := config.Read()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := algorithms.RoundRobin(cfg.Servers)

		fmt.Fprintln(w, "You are on the server", s)
	})

	fmt.Printf("Listening on http://localhost:%d\n", cfg.Port)
	log.Fatalln(http.ListenAndServe(fmt.Sprint(":", cfg.Port), nil))
}
