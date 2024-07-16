package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	key := 3
	fmt.Printf("Array: %v, Key: %d\n", slice, key)
	result := BinarySearch(slice, key)
	fmt.Println("Is key present in the array: ", result)
}

func BinarySearch(data []int, key int) bool {
	start := 0
	end := len(data) - 1

	for start <= end {
		mid := (start + end) / 2
		if data[mid] == key {
			return true
		} else if data[mid] < key {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return false
}
