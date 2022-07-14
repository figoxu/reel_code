package dp

func Fib(n int) int { // 递归
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func FibDp(n int, memories []int) int { // 递推
	if n <= 1 {
		return n
	} else if memories[n] == 0 {
		v := FibDp(n-1, memories) + FibDp(n-2, memories)
		memories[n] = v
	}
	return memories[n]
}
