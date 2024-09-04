package main

import "fmt"

func main() {
	arr := []int{3, 7, 5, 6, 1, 4, 2}
	size := len(arr)
	fmt.Println("ArrayIndexMaxDiff:", ArrayIndexMaxDiff(arr, size))
}

// ArrayIndexMaxDiff finds the maximum difference between indices (j - i),
// where j > i and arr[j] > arr[i].
func ArrayIndexMaxDiff(arr []int, size int) int {
	// Base case: if array size is less than 2, no valid pair exists
	if size < 2 {
		return -1
	}

	// Initialize arrays to store the minimum value on the left and maximum value on the right
	minLeft := make([]int, size)
	maxRight := make([]int, size)

	// Fill minLeft with the minimum values from the left
	minLeft[0] = arr[0]
	for i := 1; i < size; i++ {
		minLeft[i] = min(minLeft[i-1], arr[i])
	}

	// Fill maxRight with the maximum values from the right
	maxRight[size-1] = arr[size-1]
	for i := size - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], arr[i])
	}

	// Initialize variables to find the maximum index difference
	i, j, maxDiff := 0, 0, -1

	// Traverse both arrays to find the maximum difference where minLeft[i] < maxRight[j]
	for i < size && j < size {
		if minLeft[i] < maxRight[j] {
			maxDiff = max(maxDiff, j-i)
			j++
		} else {
			i++
		}
	}

	return maxDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
