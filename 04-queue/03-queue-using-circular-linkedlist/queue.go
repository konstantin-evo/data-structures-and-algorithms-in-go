package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type QueueLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Peek(value int) int {
	if q.IsEmpty() {
		fmt.Println("Empty queue")
		return 0
	}
	return q.tail.next.value
}

func (q *QueueLinkedList) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
		q.tail.next = q.head
	} else {
		temp.next = q.head
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *QueueLinkedList) Remove() (int, bool) {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return 0, false
	}

	var value int
	value = q.head.value
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.tail.next = q.head
	}
	q.size--
	return value, true
}

func (q *QueueLinkedList) Print() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	current := q.head
	for {
		fmt.Print(current.value, " ")
		current = current.next
		if current == q.head {
			break
		}
	}
	fmt.Println()
}
