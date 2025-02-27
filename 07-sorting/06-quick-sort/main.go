package main

import "fmt"

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)

	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

func quickSort(arr []int) []int {
	// Base case: if the arr has 0 or 1 elements, it's already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Choosing the pivot element (middle element for balance)
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var left, right []int

	// Partitioning: splitting elements into two groups
	for i, v := range arr {
		// Skip the pivot itself
		if i == pivotIndex {
			continue
		}

		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}
