package algorithms

import (
	"testing"
)

var fibs = []struct {
	in  int
	out int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{10, 55},
}

func TestFibonacci(t *testing.T) {
	for _, fib := range fibs {
		r := Fibonacci(fib.in)

		if r != fib.out {
			t.Errorf("At n=%v: Expected %v but got %v", fib.in, fib.out, r)
		}
	}
}

func TestFibonacciImproved(t *testing.T) {
	for _, fib := range fibs {
		r := FibonacciImproved(fib.in)

		if r != fib.out {
			t.Errorf("At n=%v: Expected %v but got %v", fib.in, fib.out, r)
		}
	}
}

// run all benchmarks: go test -bench .
func BenchmarkFibonacci1(b *testing.B)  { benchmarkFibonacci(1, b) }
func BenchmarkFibonacci2(b *testing.B)  { benchmarkFibonacci(2, b) }
func BenchmarkFibonacci3(b *testing.B)  { benchmarkFibonacci(3, b) }
func BenchmarkFibonacci10(b *testing.B) { benchmarkFibonacci(10, b) }
func BenchmarkFibonacci20(b *testing.B) { benchmarkFibonacci(20, b) }
func BenchmarkFibonacci40(b *testing.B) { benchmarkFibonacci(40, b) }

func benchmarkFibonacci(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(i)
	}
}

func BenchmarkFibonacciImproved1(b *testing.B)      { benchmarkFibonacciImproved(1, b) }
func BenchmarkFibonacciImproved2(b *testing.B)      { benchmarkFibonacciImproved(2, b) }
func BenchmarkFibonacciImproved3(b *testing.B)      { benchmarkFibonacciImproved(3, b) }
func BenchmarkFibonacciImproved10(b *testing.B)     { benchmarkFibonacciImproved(10, b) }
func BenchmarkFibonacciImproved20(b *testing.B)     { benchmarkFibonacciImproved(20, b) }
func BenchmarkFibonacciImproved40(b *testing.B)     { benchmarkFibonacciImproved(40, b) }
func BenchmarkFibonacciImproved100(b *testing.B)    { benchmarkFibonacciImproved(100, b) }
func BenchmarkFibonacciImproved250(b *testing.B)    { benchmarkFibonacciImproved(250, b) }
func BenchmarkFibonacciImproved500(b *testing.B)    { benchmarkFibonacciImproved(500, b) }
func BenchmarkFibonacciImproved1000(b *testing.B)   { benchmarkFibonacciImproved(1000, b) }
func BenchmarkFibonacciImproved10000(b *testing.B)  { benchmarkFibonacciImproved(10000, b) }
func BenchmarkFibonacciImproved100000(b *testing.B) { benchmarkFibonacciImproved(100000, b) }

func benchmarkFibonacciImproved(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibonacciImproved(i)
	}
}
