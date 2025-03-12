package main

func main() {
	queue := &QueueLinkedList{}

	queue.Add(10)
	queue.Add(20)
	queue.Add(30)

	queue.Print() // Output: Queue: [10 20 30]

	queue.Remove()
	queue.Print() // Output: Queue: [20 30]
}
