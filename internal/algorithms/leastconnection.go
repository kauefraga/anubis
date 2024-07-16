package algorithms

import "sync"

func LeastConnection() Algorithm {
	var serversConnectionCount map[string]int
	var m sync.Mutex

	return func(servers []string) string {
		m.Lock()
		defer m.Unlock()

		if serversConnectionCount == nil {
			serversConnectionCount = make(map[string]int, len(servers))

			for _, s := range servers {
				serversConnectionCount[s] = 0
			}
		}

		var leastConnectionServer string

		// magic number, minimum connections?
		// ref: https://www.geeksforgeeks.org/load-balancing-algorithms/#21-least-connection-method-load-balancing-algorithms
		minConnections := 1_000_000

		for server, count := range serversConnectionCount {
			if count < minConnections {
				minConnections = count
				leastConnectionServer = server
			}
		}

		if leastConnectionServer != "" {
			serversConnectionCount[leastConnectionServer] = minConnections + 1
		}

		return leastConnectionServer
	}
}
