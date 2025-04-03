package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)

	fmt.Print("5th node in Postorder Traversal: ")
	tree.NthPostOrder(1)
}

// NthPostOrder finds and prints the N-th node in post-order traversal
func (t *Tree) NthPostOrder(index int) {
	var counter int
	nthPostOrder(t.root, index, &counter)
}

// nthPostOrder performs a recursive post-order traversal to find the N-th node
func nthPostOrder(node *Node, index int, counter *int) {
	// Base case: stop if the node is nil
	if node == nil {
		return
	}

	// Recursively traverse the left and right subtrees
	nthPostOrder(node.left, index, counter)
	nthPostOrder(node.right, index, counter)

	// Increment counter after processing left and right children,
	// to visit the current node last.
	*counter++
	if *counter == index {
		fmt.Print(node.value)
		return
	}
}
