package algorithms

import (
	"sync"
	"sync/atomic"

	"github.com/kauefraga/anubis/internal/models"
)

func LeastConnection() Algorithm {
	var serversConnectionCount map[models.Servers]*int32
	var mu sync.Mutex
	var once sync.Once

	return func(servers []*models.Servers) *models.Servers {
		// ensures that will be executed one time
		once.Do(func() {
			if serversConnectionCount == nil {
				serversConnectionCount = make(map[models.Servers]*int32)

				for _, s := range servers {
					var count int32 = 0
					serversConnectionCount[*s] = &count
				}
			}
		})

		mu.Lock()
		defer mu.Unlock()

		var leastConnectionServer *models.Servers // least connection server

		// magic number, minimum connections?
		// ref: https://www.geeksforgeeks.org/load-balancing-algorithms/#21-least-connection-method-load-balancing-algorithms
		var minConnections int32 = 1_000_000

		for server, count := range serversConnectionCount {
			currenctCount := atomic.LoadInt32(count)
			if currenctCount < minConnections {
				minConnections = currenctCount
				leastConnectionServer = &server
			}
		}

		if leastConnectionServer != nil {
			atomic.AddInt32(serversConnectionCount[*leastConnectionServer], minConnections+1)
		}

		return leastConnectionServer
	}
}
