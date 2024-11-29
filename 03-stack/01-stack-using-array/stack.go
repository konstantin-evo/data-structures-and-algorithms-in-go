package main

import "fmt"

// StackInt Stack implementation using slices
type StackInt struct {
	elements []int
}

// isEmpty() function returns true if stack is empty or false in all other cases.
func (s *StackInt) IsEmpty() bool {
	length := len(s.elements)
	return length == 0
}

// length() function returns the number of elements in the stack.
func (s *StackInt) Length() int {
	length := len(s.elements)
	return length
}

// The print function will print the elements of the array.
func (s *StackInt) Print() {
	for _, value := range s.elements {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

// push() function append value to the data.
func (s *StackInt) Push(value int) {
	s.elements = append(s.elements, value)
}

/*
In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the data array and return it.
*/
func (s *StackInt) Pop() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack is empty.")
		return 0
	}
	length := len(s.elements)
	res := s.elements[length-1]
	s.elements = s.elements[:length-1]
	return res
}

/*
top() function will first check that the stack is not empty

then returns the value stored in the top element
of stack (does not remove it).
*/
func (s *StackInt) Top() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack is empty.")
		return 0
	}
	length := len(s.elements)
	res := s.elements[length-1]
	return res
}
