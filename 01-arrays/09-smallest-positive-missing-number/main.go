package main

import "fmt"

func main() {
	slice := []int{8, 5, 6, 1, 9, 11, 2, 7, 4, 10}
	result := SmallestPositiveMissingNumber(slice, len(slice))
	fmt.Println(result)
}

func SmallestPositiveMissingNumber(arr []int, size int) int {
	present := make([]bool, size+1)

	for i := 0; i < size; i++ {
		if arr[i] > 0 && arr[i] <= size {
			present[arr[i]] = true
		}
	}

	for i := 1; i <= size; i++ {
		if !present[i] {
			return i
		}
	}

	return -1
}
