package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	MaxMinArr(slice, len(slice))
	fmt.Println(slice)
}

func MaxMinArr(arr []int, size int) {
	result := make([]int, size)
	left, right := 0, size-1

	// Use a flag to switch between maximum and minimum elements
	flag := true

	for i := 0; i < size; i++ {
		if flag {
			result[i] = arr[right]
			right--
		} else {
			result[i] = arr[left]
			left++
		}
		flag = !flag
	}

	copy(arr, result)
}
