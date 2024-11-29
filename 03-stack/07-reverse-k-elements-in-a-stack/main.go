package main

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	stack.Print() // Output: Stack: [10 20 30 40]

	reverseKElementInStack(stack, 2)
	stack.Print() // Output: Stack: [10 20 40 30]
}

func reverseKElementInStack(stk *Stack, k int) {
	que := &Queue{}

	// Pop the first k elements from the stack and enqueue them
	for i := 0; i < k; i++ {
		que.Enqueue(stk.Pop())
	}

	// Dequeue elements from the queue and push them back onto the stack
	for !que.IsEmpty() {
		stk.Push(que.Dequeue().(int))
	}
}
