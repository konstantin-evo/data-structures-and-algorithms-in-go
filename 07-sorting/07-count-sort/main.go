package main

import "fmt"

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)

	sortedArr := countSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

func countSort(arr []int) []int {
	// Base case: if the arr has 0 or 1 elements, it's already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Find the max element in the array to determine the range
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Create a count array to store the count of each element
	count := make([]int, max+1)

	// Count the occurrences of each element
	for _, num := range arr {
		count[num]++
	}

	// Modify the count array to store the cumulative count
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	// Create a sorted array
	sortedArr := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		sortedArr[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	return sortedArr
}
