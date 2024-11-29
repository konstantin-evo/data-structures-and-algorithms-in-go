package main

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Print() // Output: Stack: [10 20 30]

	sortedInsert(stack, 15)
	stack.Print() // Output: Stack: [10 15 20 30]
}

func sortedInsert(s *Stack, element int) {
	if s.Top() <= element || s.IsEmpty() {
		s.Push(element)
	} else {
		// Remove the top element
		temp := s.Pop()
		// Recursive call to sortedInsert for the remaining stack
		sortedInsert(s, element)
		// Push the temporarily removed element back
		s.Push(temp)
	}
}
