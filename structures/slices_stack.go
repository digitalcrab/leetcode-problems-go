package structures

// SlicesStack represents a stack implemented with slice.
// Stack is a collection of elements with two main options:
// - push, which adds an element to the collection
// - pop, removes the most recent added element
// - (optionally) peak, which can, without modifying the stack return the value of the recent added element.
type SlicesStack []any

// Push adds an element to the collection
func (s *SlicesStack) Push(v any) {
	*s = append(*s, v)
}

// Pop removes the most recent added element
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
func (s *SlicesStack) Peek() any {
	dataLen := len(*s)
	if dataLen == 0 {
		return nil
	}
	return (*s)[dataLen-1]
}
