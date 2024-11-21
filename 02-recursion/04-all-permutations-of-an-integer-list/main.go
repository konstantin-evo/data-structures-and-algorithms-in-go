package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := []int{1, 2, 3}
	Permutation(data, 0, len(data))
}

func Permutation(data []int, i int, len int) {
	if len == i {
		temp := "{"
		for k := 0; k < len; k++ {
			temp += strconv.Itoa(data[k])
			temp += " "
		}
		temp += "}"
		fmt.Print(temp)
		return
	}

	for j := i; j < len; j++ {
		data[i], data[j] = data[j], data[i] //swap
		Permutation(data, i+1, len)
		data[i], data[j] = data[j], data[i] //swap back
	}
}
