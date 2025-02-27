package main

import "fmt"

func main() {
	arr := []int{1, 0, 1, 0, 1, 0, 1, 0}
	size := len(arr)
	fmt.Println("Array before Sort:", arr)
	swaps := Partition01(arr, size)
	fmt.Println("Sorted Array:", arr)
	fmt.Println("Minimum Swaps:", swaps)
}

func Partition01(arr []int, size int) int {
	count := 0
	left, right := 0, size-1

	for left < right {
		// Move left pointer until we find a 1
		for left < right && arr[left] == 0 {
			left++
		}
		// Move right pointer until we find a 0
		for left < right && arr[right] == 1 {
			right--
		}
		// Swap if left is still less than right
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			count++
			left++
			right--
		}
	}

	return count
}
