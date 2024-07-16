package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Array: ", slice)
	sum := sumArr(slice)
	fmt.Println("Sum: ", sum)
}

func sumArr(data []int) int {

	total := 0

	for _, value := range data {
		total += value
	}

	return total
}
