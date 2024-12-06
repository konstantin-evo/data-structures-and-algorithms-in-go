package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := LevelOrderBinaryTree(arr)
	fmt.Println("PreOrder Traversal:")
	tree.PrintBreadthFirst()
}

func (t *Tree) PrintBreadthFirst() {

	// If the tree is empty, return immediately
	if t.root == nil {
		return
	}

	// Initialize a new queue and enqueue the root node
	queue := &Queue{}
	queue.Enqueue(t.root)

	// Process nodes until the queue is empty
	for !queue.IsEmpty() {
		// Dequeue the next node and cast it to *Node
		current := queue.Dequeue().(*Node)

		// Print the value of the current node
		fmt.Print(current.value, " ")

		// Enqueue the left child if it exists
		if current.left != nil {
			queue.Enqueue(current.left)
		}

		// Enqueue the right child if it exists
		if current.right != nil {
			queue.Enqueue(current.right)
		}
	}
}
