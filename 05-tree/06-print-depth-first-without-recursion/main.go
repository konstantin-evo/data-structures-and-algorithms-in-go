package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := LevelOrderBinaryTree(arr)
	fmt.Println("DepthFirst Traversal:")
	tree.PrintDepthFirst()
}

func (t *Tree) PrintDepthFirst() {
	// If the tree is empty, return immediately
	if t.root == nil {
		return
	}

	// Initialize a new stack and push the root node
	stack := &Stack{}
	stack.Push(t.root)

	// Process nodes until the stack isn't empty
	for !stack.IsEmpty() {

		// here comment
		current := stack.Pop()
		fmt.Print(current.value, " ")

		// here comment
		if current.right != nil {
			stack.Push(current.right)
		}

		// here comment
		if current.left != nil {
			stack.Push(current.left)
		}

	}
}
