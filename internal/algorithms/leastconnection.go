package algorithms

func LeastConnection() Algorithm {
	var serversConnectionCount map[string]int

	return func(servers []string) string {
		if serversConnectionCount == nil {
			serversConnectionCount = make(map[string]int, len(servers))

			for _, s := range servers {
				serversConnectionCount[s] = 0
			}
		}

		var leastConnectionServer string // least connection server

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
