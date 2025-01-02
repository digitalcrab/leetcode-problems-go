package main

import "strconv"

type op func(a, b int) int

// just a tiny map for quicker search of operators and implementation
var ops = map[string]op{
	"-": func(a, b int) int { return a - b },
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b }, // The division between two integers always **truncates toward zero**.
}

// From wiki:
// The concept of a stack, a last-in/first-out construct, is integral to the left-to-right evaluation of RPN.
// In the example 3 4 −, first the 3 is put onto the stack, then the 4; the 4 is now on top and the 3 below it.
// The subtraction operator removes the top two items from the stack, performs 3 − 4, and puts the result
// of −1 onto the stack.

func evalRPN(tokens []string) int {
	// just one number?
	if len(tokens) == 1 {
		n, _ := strconv.Atoi(tokens[0])
		return n
	}

	// not valid expression
	if len(tokens) < 3 {
		return 0
	}

	// prepare the stack of ints
	var stack []int

	// loop over all tokens
	for _, token := range tokens {
		// if it's operator
		if fn, isOp := ops[token]; isOp {
			// pop 3 numbers from the stack
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			// truncate the stack
			stack = stack[:len(stack)-2]
			// push the result
			stack = append(stack, fn(a, b))
			continue
		}

		// if it's a number we should push to the stack
		n, _ := strconv.Atoi(token) // error we skip as based on constraint it's not possible
		stack = append(stack, n)
	}

	// at the end that should be only one number
	return stack[0]
}
