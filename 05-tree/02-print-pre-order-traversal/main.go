package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := LevelOrderBinaryTree(arr)
	fmt.Println("PreOrder Traversal:")
	tree.PrintPreOrder()
}

func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
}

func printPreOrder(n *Node) {
	// Base case: stop recursion if the node is nil
	if n == nil {
		return
	}

	fmt.Print(n.value, " ") // Print the current node
	printPreOrder(n.left)   // Recursively visit the left subtree
	printPreOrder(n.right)  // Recursively visit the right subtree
}
