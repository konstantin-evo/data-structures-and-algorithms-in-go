package main

import "fmt"

func main() {
	queue := &Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	queue.Print() // Output: Queue: [10 20 30]

	reverseQueue(queue) // Reversing the queue

	queue.Print() // Output: Queue: [30 20 10]
}

// reverseQueue reverses the given queue using a stack.
func reverseQueue(q *Queue) {

	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	s := &Stack{}

	// Dequeue all elements from the queue and push them onto the stack
	for !q.IsEmpty() {
		val, ok := q.Dequeue().(int) // Type assertion to ensure correct type
		if !ok {
			panic("Can't dequeue. The type is not correct.")
		}
		s.Push(val)
	}

	// Pop all elements from the stack and enqueue them back into the queue
	for !s.IsEmpty() {
		q.Enqueue(s.Pop())
	}
}
