package models

import (
	"sync"
)

// the only way that I found to avoid recycle import error
type Servers struct {
	Url  string
	Sync *sync.Mutex
}
