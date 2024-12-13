package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	fmt.Println("DepthFirst Traversal:")
	tree.PrintLevelOrderLineByLine()
}

func (t *Tree) PrintLevelOrderLineByLine() {
	// If the tree is empty, return immediately
	if t.root == nil {
		return
	}

	// Start by adding the root node to the current level queue
	currLevelQueue := &Queue{}
	nextLevelQueue := &Queue{}
	currLevelQueue.Enqueue(t.root)

	for !currLevelQueue.IsEmpty() {
		result := "" // String to store the values of nodes at the current level

		for !currLevelQueue.IsEmpty() {
			currentNode := currLevelQueue.Dequeue().(*Node) // Dequeue a node from the current level
			result += strconv.Itoa(currentNode.value) + " " // Append the node's value to the result string

			if currentNode.left != nil { // If the left child exists, enqueue it into the next level queue
				nextLevelQueue.Enqueue(currentNode.left)
			}
			if currentNode.right != nil { // If the right child exists, enqueue it into the next level queue
				nextLevelQueue.Enqueue(currentNode.right)
			}
		}

		fmt.Print(result + "; ")
		currLevelQueue = nextLevelQueue // Move to the next level by switching queues
		nextLevelQueue = &Queue{}       // Reset the next level queue for the upcoming level
	}
}
