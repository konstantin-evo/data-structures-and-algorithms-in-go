package main

import "fmt"

func main() {
	stack := &StackInt{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Print() // Output: Stack: [10 20 30]

	stack.Pop()
	stack.Print() // Output: Stack: [10 20]

	fmt.Println(stack.Top())     // Output: 20, true
	fmt.Println(stack.IsEmpty()) // Output: false
}
