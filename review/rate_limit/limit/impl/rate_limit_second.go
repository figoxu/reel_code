package limit

import (
	"sync/atomic"
	"time"

	"rate_limit/limit"
)

type RateLimitSecond struct {
	cfg int32
	max *int32
}

func NewRateLimitSecond(max int32) limit.RateLimit {
	out := &RateLimitSecond{
		cfg: max,
		max: &max,
	}
	go out.start()
	return out
}

func (p *RateLimitSecond) start() {
	for {
		time.Sleep(time.Second)
		atomic.CompareAndSwapInt32(p.max, atomic.LoadInt32(p.max), p.cfg)
	}
}

func (p *RateLimitSecond) Acquire() error {
	//atomic.CompareAndSwapInt32(p.max, atomic.LoadInt32(p.max), atomic.(p.max))
	return nil
}

func (p *RateLimitSecond) Release() error {
	//<-p.ch
	return nil
}
