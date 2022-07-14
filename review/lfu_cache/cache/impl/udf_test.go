package impl_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"lfu/cache"
	"lfu/cache/impl"
)

func TestImpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "RepoImpl Suite", []Reporter{reporters.NewJUnitReporter("junit.xml")})
}

var _ = Describe("UdfCacheTest", func() {
	getManyTimes := func(c cache.UdfCache, key string, times int) {
		for i := 0; i < times; i++ {
			c.Get(key)
		}
	}
	It("recycle by count", func() {
		c := impl.NewUdfCache(2)
		c.Set("one", "1")
		getManyTimes(c, "one", 3)
		c.Set("two", "2")
		getManyTimes(c, "two", 2)
		c.Set("three", "3")
		Ω(c.Get("one")).Should(Equal("1"))
		Ω(c.Get("three")).Should(Equal("3"))
		Ω(c.Get("two")).Should(BeEmpty())
	})
	It("recyle by count and ttl", func() {
		c := impl.NewUdfCache(2)
		c.Set("one", "1")
		getManyTimes(c, "two", 3)
		c.Set("two", "2")
		getManyTimes(c, "two", 3)
		c.Set("three", "3")
		Ω(c.Get("two")).Should(Equal("2"))
		Ω(c.Get("three")).Should(Equal("3"))
		Ω(c.Get("one")).Should(BeEmpty())
	})
})
