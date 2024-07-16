package algorithms

import (
	"sync"

	"github.com/kauefraga/anubis/internal/models"
)

type Algorithm func(servers []*models.Servers) *models.Servers

func RoundRobin() Algorithm {
	serverCount := 0
	var mu sync.Mutex

	return func(servers []*models.Servers) *models.Servers {
		mu.Lock()
		defer mu.Unlock()

		if serverCount >= len(servers) {
			serverCount = 0
		}

		currentServer := servers[serverCount]

		serverCount += 1

		return currentServer
	}
}
