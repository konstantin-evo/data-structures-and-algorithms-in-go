package main

import "fmt"

func main() {
	// Example array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create the binary tree
	tree := LevelOrderBinaryTree(arr)

	// Print the root value (should be 1)
	fmt.Println("Root:", tree.root.value)
}

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	// Base case: if the start index is out of bounds, return nil
	if start >= size {
		return nil
	}

	// Create a new node with the current array value
	node := &Node{value: arr[start]}

	// Recursively assign the left and right children
	node.left = levelOrderBinaryTree(arr, 2*start+1, size)
	node.right = levelOrderBinaryTree(arr, 2*start+2, size)

	return node
}
