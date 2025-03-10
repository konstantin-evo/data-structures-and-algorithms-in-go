package main

import "fmt"

func main() {
	queue := &QueueUsingStack{}

	queue.Add(10)
	queue.Add(20)
	queue.Add(30)

	queue.Print() // Output: Queue: [10 20 30]

	fmt.Println("The Peek element is ", queue.Peek())

	queue.Remove()
	queue.Print() // Output: Queue: [20 30]
}
