package main

import "fmt"

func main() {
	queue := &Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)

	queue.Print() // Output: Queue: [10 20 30 40 50]

	reverseKElementInQueue(queue, 3) // Reversing the queue

	queue.Print() // Output: Queue: [30 20 10 40 50]
}

func reverseKElementInQueue(q *Queue, k int) {

	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	if k > q.Length() || k <= 0 {
		fmt.Println("Invalid value of k")
		return
	}

	s := &Stack{}

	// Step 1: Dequeue the first k elements from the queue and push them onto the stack
	for i := 0; i < k; i++ {
		val, ok := q.Dequeue().(int) // Type assertion to ensure correct type
		if !ok {
			panic("Can't dequeue. The type is not correct.")
		}
		s.Push(val)
	}

	// Step 2: Pop elements from the stack and enqueue them back to the queue
	for i := 0; i < k; i++ {
		q.Enqueue(s.Pop())
	}

	// Step 3: Move the remaining elements to the back of the queue
	for i := 0; i < q.size-k; i++ {
		val, ok := q.Dequeue().(int)
		if !ok {
			panic("Can't dequeue. The type is not correct.")
		}
		q.Enqueue(val)
	}
}
