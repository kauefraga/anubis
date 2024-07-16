package algorithms

import (
	"sync"

	"github.com/kauefraga/anubis/internal/models"
)

func LeastConnection() Algorithm {
	var serversConnectionCount map[models.Servers]int
	var mu sync.Mutex
	var once sync.Once

	return func(servers []*models.Servers) *models.Servers {
		// ensures that will be executed one time
		once.Do(func() {
			mu.Lock()
			defer mu.Unlock()
			if serversConnectionCount == nil {
				serversConnectionCount = make(map[models.Servers]int)

				for _, s := range servers {
					serversConnectionCount[*s] = 0
				}
			}
		})

		mu.Lock()
		defer mu.Unlock()

		var leastConnectionServer *models.Servers // least connection server

		// magic number, minimum connections?
		// ref: https://www.geeksforgeeks.org/load-balancing-algorithms/#21-least-connection-method-load-balancing-algorithms
		minConnections := 1_000_000

		for server, count := range serversConnectionCount {
			if count < minConnections {
				minConnections = count
				leastConnectionServer = &server
			}
		}

		if leastConnectionServer != nil {
			serversConnectionCount[*leastConnectionServer] = minConnections + 1
		}

		return leastConnectionServer
	}
}
