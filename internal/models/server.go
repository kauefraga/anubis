package models

import (
	"sync"
)

// Avoid recycle import error (loadbalancer <-> algorithms)
type Server struct {
	Url  string
	Sync *sync.Mutex
}
