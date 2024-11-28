package main

import "fmt"

func main() {
	expr := "{[()()]}"
	if IsBalancedParenthesis(expr) {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not Balanced")
	}
}

func IsBalancedParenthesis(expn string) bool {
	// Create a new stack to keep track of opening parentheses
	stack := &Stack{}

	// Iterate over each character in the expression
	for _, char := range expn {
		// If it's an opening parenthesis, push it onto the stack
		switch char {
		case '{', '(', '[':
			stack.Push(char)
		// If it's a closing parenthesis, check for a matching opening parenthesis
		case '}':
			// Pop the stack and check if the top matches the opening parenthesis
			if stack.IsEmpty() || stack.Pop() != '{' {
				return false
			}
		case ')':
			// Pop the stack and check if the top matches the opening parenthesis
			if stack.IsEmpty() || stack.Pop() != '(' {
				return false
			}
		case ']':
			// Pop the stack and check if the top matches the opening parenthesis
			if stack.IsEmpty() || stack.Pop() != '[' {
				return false
			}
		}
	}

	// If the stack is empty, all parentheses are matched and balanced
	return stack.IsEmpty()
}
