package structures

import (
	"bytes"
	"fmt"
)

// SlicesStack represents a stack implemented with slice.
// Stack is a collection of elements with two main options:
// - push, which adds an element to the collection
// - pop, removes the most recent added element
// - (optionally) peak, which can, without modifying the stack return the value of the recent added element.
type SlicesStack []any

func (s *SlicesStack) Display() string {
	buf := new(bytes.Buffer)
	buf.WriteString("[")
	for i := len(*s) - 1; i >= 0; i-- {
		_, _ = fmt.Fprintf(buf, "%v", (*s)[i])
		if i != 0 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// Push adds an element to the collection
// Time Complexity:
// O(1) on average as it simply adds one element to the slice
// O(n) worst case it it needs to grow the slice and copy over all elements
func (s *SlicesStack) Push(v any) {
	*s = append(*s, v)
}

// Pop removes the most recent added element
// Time Complexity:
// O(1) resetting the stack simply removes element and does not create a new array under the hood just creates a new
// header structure that points to the less elements
func (s *SlicesStack) Pop() any {
	dataLen := len(*s)
	if dataLen == 0 {
		return nil
	}

	// get the last element
	element := (*s)[dataLen-1]

	// reset the stack
	*s = (*s)[:dataLen-1]

	return element
}

// Peek returns the recent element without modifying the stack
// Time Complexity:
// O(1) simply returning element
func (s *SlicesStack) Peek() any {
	dataLen := len(*s)
	if dataLen == 0 {
		return nil
	}
	return (*s)[dataLen-1]
}
