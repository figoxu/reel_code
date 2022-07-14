package limit

type RateLimit interface {
	Acquire() error
	Release() error
}
