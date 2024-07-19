package algorithms

import (
	"sync"

	"github.com/kauefraga/anubis/internal/models"
)

func LeastConnection(servers []*models.Server) Algorithm {
	serversConnectionCount := make(map[models.Server]int, len(servers))
	var mu sync.Mutex

	for _, s := range servers {
		serversConnectionCount[*s] = 0
	}

	return func() *models.Server {
		mu.Lock()
		defer mu.Unlock()

		var leastConnectionServer *models.Server

		// magic number, minimum connections?
		// ref: https://www.geeksforgeeks.org/load-balancing-algorithms/#21-least-connection-method-load-balancing-algorithms
		var minConnections int = 1_000_000

		for server, count := range serversConnectionCount {
			if count < minConnections {
				minConnections = count
				leastConnectionServer = &server
			}
		}

		if leastConnectionServer != nil {
			serversConnectionCount[*leastConnectionServer]++
		}

		return leastConnectionServer
	}
}
