package main

import "fmt"

func main() {
	data := []int{1, 3, 5, 7, 9, 11, 13}
	value := 7

	if BinarySearchRecursive(data, value) {
		fmt.Printf("Value %d found in the array.\n", value)
	} else {
		fmt.Printf("Value %d not found in the array.\n", value)
	}
}

func BinarySearchRecursive(data []int, value int) bool {
	size := len(data)
	return BinarySearchRecursiveUtil(data, 0, size-1, value)
}

func BinarySearchRecursiveUtil(data []int, low int, high int, value int) bool {

	if low > high {
		return false
	}

	mid := low + (high-low)/2

	if data[mid] == value {
		return true
	}

	if data[mid] < value {
		BinarySearchRecursiveUtil(data, low, mid-1, value)
	}

	return BinarySearchRecursiveUtil(data, mid+1, high, value)
}
