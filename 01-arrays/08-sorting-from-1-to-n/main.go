package main

import "fmt"

func main() {
	slice := []int{8, 5, 6, 1, 9, 3, 2, 7, 4, 10}
	Sort1toN(slice, len(slice))
	fmt.Println(slice)
}

func Sort1toN(arr []int, size int) {
	for i := 0; i < size; i++ {
		for arr[i] != i+1 && arr[i] > 1 {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}
	}
}
