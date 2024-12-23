package main

import "fmt"

func main() {
	// Example array
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:", arr)

	// Sort in ascending order
	InsertedSort(arr, Ascending)
	fmt.Println("Sorted in ascending order:", arr)

	// Reset the array
	arr = []int{64, 34, 25, 12, 22, 11, 90}

	// Sort in descending order
	InsertedSort(arr, Descending)
	fmt.Println("Sorted in descending order:", arr)
}

func InsertedSort(arr []int, comp func(int, int) bool) {
	// Iterate through the array starting from the second element
	for i := 1; i < len(arr); i++ {
		current := arr[i] // Store the current element to be inserted
		j := i - 1        // Start comparing with the last element of the sorted portion

		// Move elements of the sorted portion that are greater (or smaller, depending on `comp`)
		// than `current` one position to the right
		for j >= 0 && comp(current, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}

		// Insert the current element into its correct position in the sorted portion
		arr[j+1] = current
	}
}

func Ascending(a, b int) bool {
	return a < b
}

func Descending(a, b int) bool {
	return a > b
}
