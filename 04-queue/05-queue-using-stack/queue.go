package main

import "fmt"

type QueueUsingStack struct {
	inputStack  Stack
	outputStack Stack
}

func (q *QueueUsingStack) Add(value int) {
	q.inputStack.Push(value)
}

func (q *QueueUsingStack) Remove() int {
	if q.outputStack.IsEmpty() {
		for !q.inputStack.IsEmpty() {
			elem := q.inputStack.Pop()
			q.outputStack.Push(elem)
		}
	}
	return q.outputStack.Pop()
}

func (q *QueueUsingStack) Peek() int {
	if q.outputStack.IsEmpty() {
		for !q.inputStack.IsEmpty() {
			elem := q.inputStack.Pop()
			q.outputStack.Push(elem)
		}
	}

	return q.outputStack.Peek()
}

func (q *QueueUsingStack) Length() int {
	return q.inputStack.Length() + q.outputStack.Length()
}

func (q *QueueUsingStack) IsEmpty() bool {
	return q.outputStack.IsEmpty() && q.inputStack.IsEmpty()
}

func (q *QueueUsingStack) Print() {
	// Check if both Stacks are empty
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	// Create a temporary slice to store inputStack elements in the correct order
	// Reverse the order of inputStack elements to match FIFO behaviour
	var temp []int
	for _, val := range q.inputStack.elements {
		temp = append(temp, val)
	}

	// Print elements from outputStack (FIFO order) first, then inputStack (reversed)
	fmt.Print("Queue: ")
	q.outputStack.Print()
	for _, val := range temp {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
