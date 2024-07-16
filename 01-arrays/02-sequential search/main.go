package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 8, 7, 6}
	key := 3
	fmt.Printf("Array: %v, Key: %d\n", slice, key)
	result := SequentialSearch(slice, key)
	fmt.Println("Is key present in the array: ", result)
}

func SequentialSearch(data []int, key int) bool {
	for _, value := range data {
		if value == key {
			return true
		}
	}

	return false
}
