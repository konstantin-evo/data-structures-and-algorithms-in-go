package main

import "fmt"

func main() {
	slice := []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	sum := MaxSubArraySum(slice)
	fmt.Printf("Array: %v, Max Sum: %d/n", slice, sum)
}

func MaxSubArraySum(array []int) int {
	if len(array) == 0 {
		return 0
	}

	maxSoFar := array[0]
	maxEndingHere := array[0]

	for i := 1; i < len(array); i++ {
		maxEndingHere = max(array[i], maxEndingHere+array[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
