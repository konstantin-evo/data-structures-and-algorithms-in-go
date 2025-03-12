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

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() int {
	length := len(s.elements)
	res := s.elements[length-1]
	s.elements = s.elements[:length-1]
	return res
}

func (s *Stack) Peek() int {
	length := len(s.elements)
	res := s.elements[length-1]
	return res
}

func (s *Stack) Print() {
	length := len(s.elements)
	for i := 0; i < length; i++ {
		fmt.Print(s.elements[i], " ")
	}
	fmt.Println()
}
