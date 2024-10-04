package main

import "fmt"

func main() {
	n := 10
	fmt.Printf("Fibonacci(%d) = %d\n", n, fibonacci(n))
}

// a time complexity of O(2^n)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// a time complexity of O(n)
func fibonacciWithoutRecursion(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
