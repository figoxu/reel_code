package cache

type UdfCache interface {
	Get(k string) string
	Set(k, v string)
}
