package inference

import (
	"sync"
	"time"
)

type Vistor struct {
	lastSeem time.Time
	tokens   int64
}

type RateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*Vistor
	rate     time.Duration
	burst    int64
}
