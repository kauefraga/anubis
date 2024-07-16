package algorithms

import "sync"

type Algorithm func(servers []string) string

func RoundRobin() Algorithm {
	serverCount := 0
	var m sync.Mutex

	return func(servers []string) string {
		m.Lock()
		defer m.Unlock()

		if serverCount >= len(servers) {
			serverCount = 0
		}

		currentServer := servers[serverCount]
		serverCount += 1

		return currentServer
	}
}
