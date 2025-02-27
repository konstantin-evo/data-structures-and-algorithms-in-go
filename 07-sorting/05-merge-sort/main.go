package main

import "fmt"

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from left and right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
