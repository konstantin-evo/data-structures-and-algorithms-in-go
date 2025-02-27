package main

import "fmt"

func main() {
	// Example array
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:", arr)

	// Sort in ascending order
	BubbleSort(arr, Ascending)
	fmt.Println("Sorted in ascending order:", arr)

	// Reset the array
	arr = []int{64, 34, 25, 12, 22, 11, 90}

	// Sort in descending order
	BubbleSort(arr, Descending)
	fmt.Println("Sorted in descending order:", arr)
}

// BubbleSort sorts an array using the Bubble Sort algorithm and a comparison function.
func BubbleSort(arr []int, comp func(int, int) bool) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no elements were swapped, the array is already sorted
		if !swapped {
			break
		}
	}

}

func Ascending(a, b int) bool {
	return a > b
}

func Descending(a, b int) bool {
	return a < b
}
