package main

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Print() // Output: Stack: [10 20 30]

	bottomInsert(stack, 15)
	stack.Print() // Output: Stack: [15 10 20 30]]
}

func bottomInsert(s *Stack, element int) {
	// Base case: If the stack is empty, push the element directly to the stack
	if s.IsEmpty() {
		s.Push(element)
	} else {
		// Recursive case: Pop the top element and hold it temporarily
		temp := s.Pop()

		// Recursively call bottomInsert to insert the element at the bottom of the stack
		bottomInsert(s, element)

		// Once the element is inserted at the bottom, push the temporarily held element back onto the stack
		s.Push(temp)
	}
}
