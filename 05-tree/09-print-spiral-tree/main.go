package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	tree.PrintSpiralTree()
}

func (t *Tree) PrintSpiralTree() {
	if t.root == nil {
		return
	}

	// Two stacks to store nodes at different levels in alternate orders
	leftToRight := new(Stack) // Stack for left-to-right traversal
	rightToLeft := new(Stack) // Stack for right-to-left traversal

	// Start with the root in the leftToRight stack
	leftToRight.Push(t.root)

	firstIteration := true // Flag to handle formatting of output

	for !leftToRight.IsEmpty() || !rightToLeft.IsEmpty() {
		var levelOutput []int

		// Process nodes in leftToRight stack and push their children to rightToLeft stack
		for !leftToRight.IsEmpty() {
			node := leftToRight.Pop()
			levelOutput = append(levelOutput, node.value)
			// Push left child first, then right child
			if node.left != nil {
				rightToLeft.Push(node.left)
			}
			if node.right != nil {
				rightToLeft.Push(node.right)
			}
		}

		// Print the current level's values if any nodes were processed
		if len(levelOutput) > 0 {
			if !firstIteration {
				fmt.Printf(" ; ")
			}
			fmt.Print(levelOutput)
			firstIteration = false
		}

		levelOutput = nil // Reset level output for next level

		// Process nodes in rightToLeft stack and push their children to leftToRight stack
		for !rightToLeft.IsEmpty() {
			node := rightToLeft.Pop()
			levelOutput = append(levelOutput, node.value)
			// Push right child first, then left child
			if node.right != nil {
				leftToRight.Push(node.right)
			}
			if node.left != nil {
				leftToRight.Push(node.left)
			}
		}

		// Print the current level's values if any nodes were processed
		if len(levelOutput) > 0 {
			fmt.Print(" ; ", levelOutput)
		}
	}

	fmt.Print(" ; ")
}
