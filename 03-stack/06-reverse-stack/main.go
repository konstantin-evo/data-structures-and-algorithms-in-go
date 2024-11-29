package main

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Print() // Output: Stack: [10 20 30]

	reverseStack(stack)
	stack.Print() // Output: Stack: [30 20 10]
}

func reverseStack(s *Stack) {
	// Base case: If the stack is empty, stop the recursion.
	if s.IsEmpty() {
		return
	}

	temp := s.Pop()       // Step 1: Remove the top element
	reverseStack(s)       // Step 2: Recursively reverse the remaining stack
	bottomInsert(s, temp) // Step 3: Insert the removed element at the bottom
}
