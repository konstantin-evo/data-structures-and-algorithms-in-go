package main

import "fmt"

func main() {
	// Example array
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:", arr)

	// Sort in ascending order
	ModifiedBubbleSort(arr, Ascending)
	fmt.Println("Sorted in ascending order:", arr)

	// Reset the array
	arr = []int{64, 34, 25, 12, 22, 11, 90}

	// Sort in descending order
	ModifiedBubbleSort(arr, Descending)
	fmt.Println("Sorted in descending order:", arr)
}

// ModifiedBubbleSort sorts an array using a modified Bubble Sort algorithm.
func ModifiedBubbleSort(arr []int, comp func(int, int) bool) {
	n := len(arr)

	for {
		// Track the last index where a swap occurred
		swapped := false
		lastSwappedIndex := 0

		for j := 0; j < n-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
				lastSwappedIndex = j + 1
			}
		}
		// If no elements were swapped, the array is already sorted
		if !swapped {
			break
		}

		// Reduce the effective size of the array
		n = lastSwappedIndex
	}

}

func Ascending(a, b int) bool {
	return a > b
}

func Descending(a, b int) bool {
	return a < b
}
