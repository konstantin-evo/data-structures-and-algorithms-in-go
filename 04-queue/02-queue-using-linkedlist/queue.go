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

func (q *QueueLinkedList) Peek() int {
	if q.IsEmpty() {
		fmt.Println("Empty queue")
		return 0
	}
	return q.head.value
}

func (q *QueueLinkedList) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *QueueLinkedList) Remove() (int, bool) {
	if q.IsEmpty() {
		fmt.Println("Empty queue")
		return 0, false
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value, true
}

func (q *QueueLinkedList) Print() {
	temp := q.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}
