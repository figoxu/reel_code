package dp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"leet/dp"
)

var _ = FDescribe("fib", func() {
	It("recursive", func() {
		Ω(dp.Fib(0)).Should(BeEquivalentTo(0))
		Ω(dp.Fib(1)).Should(BeEquivalentTo(1))
		Ω(dp.Fib(2)).Should(BeEquivalentTo(1))
		Ω(dp.Fib(3)).Should(BeEquivalentTo(2))
		Ω(dp.Fib(4)).Should(BeEquivalentTo(3))
		Ω(dp.Fib(5)).Should(BeEquivalentTo(5))
		Ω(dp.Fib(6)).Should(BeEquivalentTo(8))
	})
	It("dp", func() {
		memory := make([]int, 100)
		Ω(dp.FibDp(0, memory)).Should(BeEquivalentTo(0))
		Ω(dp.FibDp(1, memory)).Should(BeEquivalentTo(1))
		Ω(dp.FibDp(2, memory)).Should(BeEquivalentTo(1))
		Ω(dp.FibDp(3, memory)).Should(BeEquivalentTo(2))
		Ω(dp.FibDp(4, memory)).Should(BeEquivalentTo(3))
		Ω(dp.FibDp(5, memory)).Should(BeEquivalentTo(5))
		Ω(dp.FibDp(6, memory)).Should(BeEquivalentTo(8))
	})
})
