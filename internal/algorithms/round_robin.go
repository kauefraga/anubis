package algorithms

var serverCount = 0

func RoundRobin(servers []string) string {
	if serverCount >= len(servers) {
		serverCount = 0
	}

	currentServer := servers[serverCount]
	serverCount += 1

	return currentServer
}
