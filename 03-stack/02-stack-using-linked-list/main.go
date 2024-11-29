package main

import "fmt"

func main() {
	s := new(StackLinkedList)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, _ := s.Peek()
	fmt.Println("Peek() of a stack is:", val)
	fmt.Print("Stack consist following elements: ")
	for s.IsEmpty() == false {
		val, _ = s.Pop()
		fmt.Print(val, " ")
	}
}
