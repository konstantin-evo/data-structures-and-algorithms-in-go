package main

import "fmt"

func main() {
	m := 252
	n := 105

	fmt.Println("The GCD of ", m, " and ", n, " is: ", GCD(m, n))
}

func GCD(m int, n int) int {
	if n == 0 {
		return m
	}

	return GCD(n, m%n)
}
