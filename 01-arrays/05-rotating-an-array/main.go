package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	k := 2
	RotateArrayLeft(slice, k)
	fmt.Println(slice)

	slice = []int{1, 2, 3, 4, 5}
	RotateArrayRight(slice, k)
	fmt.Println(slice)
}

func RotateArrayLeft(data []int, k int) {
	n := len(data)
	if n == 0 || k%n == 0 {
		return
	}

	k = k % n
	reverse(data, 0, n-1)
	reverse(data, 0, n-k-1)
	reverse(data, n-k, n-1)
}

func RotateArrayRight(data []int, k int) {
	n := len(data)
	if n == 0 || k%n == 0 {
		return
	}

	k = k % n
	reverse(data, 0, n-1)
	reverse(data, 0, k-1)
	reverse(data, k, n-1)
}

// reverse reverses the elements of the array 'data' from index 'start' to 'end'
func reverse(data []int, start, end int) {
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}
}
