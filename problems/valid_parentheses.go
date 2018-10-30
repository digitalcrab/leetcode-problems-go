package problems

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
// - Open brackets must be closed by the same type of brackets.
// - Open brackets must be closed in the correct order.
//
// Note that an empty string is also considered valid.
//
// Time complexity : O(n) because we simply traverse the given string one character at a time
// and push and pop operations on a stack take O(1) time.
// Space complexity : O(n) as we push all opening brackets onto the stack and in the worst case,
// we will end up pushing all the brackets onto the stack.
//
func IsValidParentheses(s string) bool {
	if s == "" {
		return true
	}

	brackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	// some sort of stack
	// todo: we can use another stack implementation to save some space, but in a worst case it would be O(n) anyways
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		// if it's open one
		if closed, ok := brackets[c]; ok {
			// push the closed one as the next expected
			stack = append(stack, closed)
			continue
		}

		// nothing in the stack
		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != c {
			return false
		}

		// pop
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
