package algorithms

import (
	"sync/atomic"

	"github.com/kauefraga/anubis/internal/models"
)

type Algorithm func(servers []*models.Servers) *models.Servers

func RoundRobin() Algorithm {
	var serverCount int32 = 0

	return func(servers []*models.Servers) *models.Servers {
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
