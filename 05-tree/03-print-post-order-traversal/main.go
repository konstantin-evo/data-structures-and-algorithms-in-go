package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.PrintPostOrder()
}

func (t *Tree) PrintPostOrder() {
	fmt.Print("Post Order : ")
	printPostOrder(t.root)
	fmt.Println()
}

func printPostOrder(n *Node) {
	// Base case: if the node is nil, return (end recursion)
	if n == nil {
		return
	}

	printPostOrder(n.left)  // Recursively traverse the left subtree
	printPostOrder(n.right) // Recursively traverse the right subtree
	fmt.Print(n.value, " ") // Visit the current node (print its value)
}
