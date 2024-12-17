package main

import "fmt"

func (t *Tree) PrintInOrder() {
	fmt.Print("In Order : ")
	printInOrder(t.root)
	fmt.Println()
}

func printInOrder(n *Node) {
	// Base case: if the node is nil, return (end recursion)
	if n == nil {
		return
	}

	printInOrder(n.left)    // Recursively traverse the left subtree
	fmt.Print(n.value, " ") // Visit the current node (print its value)
	printInOrder(n.right)   // Recursively traverse the right subtree

}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.PrintInOrder()
}
