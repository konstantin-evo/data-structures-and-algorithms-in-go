package main

func main() {
	queue := &Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	queue.Print() // Output: Queue: [10 20 30]

	queue.Dequeue()
	queue.Print() // Output: Queue: [20 30]
}
