package main

import "fmt"

var parentheses = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	// empty is always valid
	if s == "" {
		return true
	}

	// can not be valid of only one char
	if len(s) == 1 {
		return false
	}

	// we're going to use stack to store the expected parentheses
	var stack []byte

	for i := range s { // Time: O(n) where n is number of chars in the string
		// get the char
		ch := s[i]

		// check if that's the open, then add to stack, what is expected next
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, parentheses[ch])
		} else {
			ln := len(stack)

			// we have nothing to expect now?
			if ln == 0 {
				return false
			}

			// for close parentheses we need to pop the stack and compare if that the one we expect
			expected := stack[ln-1]
			stack = stack[:ln-1]

			// not what we expect
			if ch != expected {
				return false
			}
		}
	}

	// it's valid only of everything consumed from the stack
	return len(stack) == 0
}

func main() {
	fmt.Printf("Parentheses sequense ()[]{} is valud: %t\n", isValid("()[]{}"))
	fmt.Printf("Parentheses sequense ([]) is valud: %t\n", isValid("([])"))
	fmt.Printf("Parentheses sequense (] is valud: %t\n", isValid("(]"))
}
