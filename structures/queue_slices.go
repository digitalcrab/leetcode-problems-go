package structures

import (
	"bytes"
	"fmt"
)

// SlicesQueue represents a queue: collection of items that maintain the order. Structure
// can be modified by adding to the back of a queue (Enqueue) and removing from front (Dequeue),
// so-called first-in-first-out (FIFO) data structure.
type SlicesQueue []any

func (s *SlicesQueue) Display() string {
	buf := new(bytes.Buffer)
	buf.WriteString("[")
	for i := len(*s) - 1; i >= 0; i-- {
		_, _ = fmt.Fprintf(buf, "%v", (*s)[i])
		if i != 0 {
			buf.WriteString("->")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// Enqueue adds value to the back of a queue
// Time Complexity:
// O(1) on average as it simply adds one element to the slice
// O(n) worst case it it needs to grow the slice and copy over all elements
func (s *SlicesQueue) Enqueue(v any) {
	*s = append(*s, v) // keep element at the back of the slice
}

// Dequeue removes value from the beginning of a queue
// Time Complexity:
// O(1) as it requires updating slice header and some other small operations
func (s *SlicesQueue) Dequeue() any {
	dataLen := len(*s)
	if dataLen == 0 {
		return nil
	}

	// get the fist element
	element := (*s)[0] // O(1)

	// reset the queue header to point to the second element
	*s = (*s)[1:] // O(1)

	return element
}
