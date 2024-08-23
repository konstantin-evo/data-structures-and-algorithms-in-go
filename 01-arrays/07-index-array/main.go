package main

import "fmt"

func main() {
	slice := []int{-1, 1, 2, 3, 4, -1, 6, 7, 8, 9}
	indexArray(slice, len(slice))
	fmt.Println(slice)
}

func indexArray(arr []int, size int) {
	result := make([]int, size)
	for i := range result {
		result[i] = -1
	}

	for _, value := range arr {
		if value != -1 {
			result[value] = value
		}
	}

	copy(arr, result)
}
