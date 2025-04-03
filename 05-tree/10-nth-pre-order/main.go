package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := LevelOrderBinaryTree(arr)

	fmt.Print("3rd node in Preorder Traversal: ")
	tree.NthPreOrder(3) // Should print 4
}

// NthPreOrder finds and prints the N-th node in pre-order traversal
func (t *Tree) NthPreOrder(index int) {
	var counter int
	nthPreOrder(t.root, index, &counter)
}

// nthPreOrder performs a recursive pre-order traversal to find the N-th node
func nthPreOrder(node *Node, index int, counter *int) {
	// Base case: stop if the node is nil
	if node == nil {
		return
	}

	*counter++
	if *counter == index {
		fmt.Print(node.value)
		return
	}

	// Recursively traverse the left and right subtrees
	nthPreOrder(node.left, index, counter)
	nthPreOrder(node.right, index, counter)
}
