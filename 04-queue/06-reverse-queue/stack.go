package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) IsEmpty() bool {
	length := len(s.elements)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.elements)
	return length
}

func (s *Stack) Print() {
	length := len(s.elements)
	for i := 0; i < length; i++ {
		fmt.Print(s.elements[i], " ")
	}
	fmt.Println()
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack in empty.")
		return 0
	}
	length := len(s.elements)
	res := s.elements[length-1]
	s.elements = s.elements[:length-1]
	return res
}

func (s *Stack) Top() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack in empty.")
		return 0
	}
	length := len(s.elements)
	res := s.elements[length-1]
	return res
}
