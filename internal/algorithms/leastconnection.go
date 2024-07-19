package algorithms

import (
	"sync"

	"github.com/kauefraga/anubis/internal/models"
)

func LeastConnection(servers []*models.Server) Algorithm {
	serversConnectionCount := make(map[models.Server]int, len(servers))
	var m sync.Mutex

	for _, s := range servers {
		serversConnectionCount[*s] = 0
	}

	return func() *models.Server {
		var leastConnectionServer *models.Server
		minConnections := 2_000_000_000 // magic number, minimum connections

		m.Lock()
		defer m.Unlock()

		for server, count := range serversConnectionCount {
			if count < minConnections {
				minConnections = count
				leastConnectionServer = &server
			}
		}

		if leastConnectionServer != nil {
			serversConnectionCount[*leastConnectionServer] += 1
		}

		return leastConnectionServer
	}
}
