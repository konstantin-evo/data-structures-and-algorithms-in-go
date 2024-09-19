package main

import "fmt"

func main() {
	number := 129
	PrintInt(number) // Outputs: 5F
	fmt.Println()
}

func PrintInt(number int) {
	// If the number is 0, stop recursion
	if number == 0 {
		return
	}

	// Recursion: divide the number by 16
	PrintInt(number / 16)

	// Get the reminder to print the corresponding hexadecimal character
	remainder := number % 16

	if remainder < 10 {
		fmt.Print(remainder)
	} else {
		fmt.Printf("%c", remainder-10+'A')
	}
}
