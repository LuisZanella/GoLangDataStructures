package main

func climbStairs(n int) int {
	return fibonacci(map[int]int{}, n)
}
func fibonacci(cache map[int]int, n int) int {
	val, ok := cache[n]
	if ok {
		return val
	}
	if n <= 1 {
		cache[n] = 1
		return cache[n]
	}
	cache[n] = fibonacci(cache, n-1) + fibonacci(cache, n-2)
	return cache[n]
}
