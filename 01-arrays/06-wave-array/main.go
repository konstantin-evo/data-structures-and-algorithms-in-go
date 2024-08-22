package main

import "fmt"

func main() {
	slice := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}
	WaveArray(slice)
	fmt.Println(slice)
}

func WaveArray(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if i%2 == 0 && arr[i] < arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		} else if i%2 != 0 && arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}
