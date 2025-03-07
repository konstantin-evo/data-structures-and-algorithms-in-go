package main

import "fmt"

func main() {
	deque := &Deque{}

	// Test adding elements to the front
	deque.AddFront(10)
	deque.AddFront(20)
	deque.AddFront(30)
	fmt.Print("After AddFront(10, 20, 30): ")
	deque.Print() // Expected: 30 20 10

	// Test adding elements to the back
	deque.AddBack(40)
	deque.AddBack(50)
	fmt.Print("After AddBack(40, 50): ")
	deque.Print() // Expected: 30 20 10 40 50

	// Test PeekFront and PeekBack
	fmt.Println("PeekFront():", deque.PeekFront()) // Expected: 30
	fmt.Println("PeekBack():", deque.PeekBack())   // Expected: 50

	// Test removing from the front
	val, ok := deque.RemoveFront()
	if ok {
		fmt.Println("RemoveFront():", val) // Expected: 30
	} else {
		fmt.Println("RemoveFront(): Queue is empty")
	}
	deque.Print() // Expected: 20 10 40 50

	// Test removing from the back
	val, ok = deque.RemoveBack()
	if ok {
		fmt.Println("RemoveBack():", val) // Expected: 50
	} else {
		fmt.Println("RemoveBack(): Queue is empty")
	}
	deque.Print() // Expected: 20 10 40

	// Test removing all elements
	for i := 0; i < 3; i++ {
		val, ok = deque.RemoveFront()
		if ok {
			fmt.Println("RemoveFront():", val)
		} else {
			fmt.Println("RemoveFront(): Queue is empty")
		}
	}
	deque.Print() // Expected: (empty queue)

	// Test operations on an empty deque
	val, ok = deque.RemoveFront()
	fmt.Println("RemoveFront():", val, "Success:", ok) // Expected: 0, false (QueueEmptyException)

	val, ok = deque.RemoveBack()
	fmt.Println("RemoveBack():", val, "Success:", ok) // Expected: 0, false (QueueEmptyException)

	deque.Print() // Expected: (empty queue)

	// Test adding after emptying
	deque.AddFront(100)
	deque.AddBack(200)
	fmt.Print("After AddFront(100) and AddBack(200): ")
	deque.Print() // Expected: 100 200
}
