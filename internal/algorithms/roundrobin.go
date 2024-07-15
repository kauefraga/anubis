package algorithms

type Algorithm func(servers []string) string

func RoundRobin() Algorithm {
	serverCount := 0

	return func(servers []string) string {
		if serverCount >= len(servers) {
			serverCount = 0
		}

		currentServer := servers[serverCount]
		serverCount += 1

		return currentServer
	}
}
