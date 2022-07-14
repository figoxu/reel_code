package limit

import "rate_limit/limit"

type RateLimitConcurrent struct {
	ch chan interface{}
}

func NewRateLimitConcurrent(max int) limit.RateLimit {
	return &RateLimitConcurrent{
		ch: make(chan interface{}, max),
	}
}

func (p *RateLimitConcurrent) Acquire() error {
	p.ch <- nil
	return nil
}

func (p *RateLimitConcurrent) Release() error {
	<-p.ch
	return nil
}
