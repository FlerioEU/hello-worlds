package algorithms

// Calculates the Fibonacci number of given integer recursively
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}

// Calculates the Fibonacci number of given integer but it is
// not using a recursive method but uses an array instead
func FibonacciImproved(n int) int {
	fibs := []int{0, 1}

	for i := 2; i <= n; i++ {
		fib := fibs[i-2] + fibs[i-1]
		fibs = append(fibs, fib)
	}

	return fibs[n]
}
