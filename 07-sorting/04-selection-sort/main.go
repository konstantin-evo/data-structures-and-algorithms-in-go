package main

import "fmt"

func main() {
	// Example array
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:", arr)

	// Sort in ascending order
	SelectingSort(arr)
	fmt.Println("Sorted in ascending order:", arr)

}

func SelectingSort(arr []int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		// Assume the first element is the smallest
		min := i
		// Find the minimum element in unsorted part of array
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min] {

				min = j
			}
		}
		// Swap the found minimum element with the first element of the unsorted part
		arr[i], arr[min] = arr[min], arr[i]
	}
}
