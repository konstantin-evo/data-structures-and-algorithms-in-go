package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 0, 1, 2, 0}
	size := len(arr)
	fmt.Println("Array before Sort:", arr)
	Partition012(arr, size)
	fmt.Println("Sorted Array:", arr)
}

func Partition012(arr []int, size int) {
	low, mid, high := 0, 0, size-1

	for mid <= high {
		switch arr[mid] {
		case 0: // Если 0, меняем с элементом в `low` и двигаем `low` и `mid`
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		case 1: // Если 1, просто двигаем `mid`
			mid++
		case 2: // Если 2, меняем с `high` и двигаем `high` (но `mid` не двигаем сразу)
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}
