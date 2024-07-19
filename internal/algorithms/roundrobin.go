package algorithms

import (
	"sync/atomic"

	"github.com/kauefraga/anubis/internal/models"
)

type Algorithm func() *models.Server

func RoundRobin(servers []*models.Server) Algorithm {
	var serverCount int32 = 0

	return func() *models.Server {
		// Will loop until the function success
		for {
			current := atomic.LoadInt32(&serverCount)
			next := current + 1
			if next >= int32(len(servers)) {
				next = 0
			}

			if atomic.CompareAndSwapInt32(&serverCount, current, next) {
				return servers[current]
			}
		}
	}
}
