package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	fmt.Println("LevelOrder Traversal:")
	tree.PrintLevelOrderLineByLine2()
}

func (t *Tree) PrintLevelOrderLineByLine2() {
	// If the tree is empty, return immediately
	if t.root == nil {
		return
	}

	queue := &Queue{}     // Create a queue to hold nodes for level-order traversal
	queue.Enqueue(t.root) // Enqueue the root node as the first element

	// While there are nodes in the queue, continue processing
	for queue.Length() > 0 {
		nodesAtCurrentLevel := queue.Length()

		for nodesAtCurrentLevel > 0 {
			// Dequeue the node at the front of the queue
			currentNode := queue.Dequeue().(*Node)
			// Print the current node value
			fmt.Printf("%d ", currentNode.value)

			// Enqueue children if they exist
			if currentNode.left != nil {
				queue.Enqueue(currentNode.left)
			}
			if currentNode.right != nil {
				queue.Enqueue(currentNode.right)
			}

			// Decrease the count of nodes at the current level
			nodesAtCurrentLevel--
		}
		// Print a separator after each level
		fmt.Print("; ")
	}
}
