package main

import "fmt"

func main() {
	num := 3
	towerOfHanoi(num, 'A', 'C', 'B')
}

func towerOfHanoi(num int, src byte, dst byte, temp byte) {
	if num == 1 {
		return
	}

	// Move the top (num-1) disks from src to temp using dst as a helper
	towerOfHanoi(num-1, src, temp, dst)

	// Move the last disk directly from src to dst
	fmt.Printf("Move %d disk from peg %c to peg %c\n", num, src, dst)

	// Move the (num-1) disks from temp to dst using src as a helper
	towerOfHanoi(num-1, temp, dst, src)
}
