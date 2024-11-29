package main

import "fmt"

func main() {
	s := &Stack{}
	s.Push(3)
	s.Push(1)
	s.Push(4)
	s.Push(2)

	fmt.Println("Original Stack:")
	s.Print()

	sortStack(s)

	fmt.Println("Sorted Stack:")
	s.Print()
}

func sortStack(s *Stack) {
	if s.IsEmpty() {
		return
	}

	// Remove the top element
	temp := s.Pop()
	// Recursively sort the remaining stack
	sortStack(s)
	// Insert the removed element back in sorted order
	sortedInsert(s, temp)
}
