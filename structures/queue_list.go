package structures

import (
	"bytes"
	"container/list"
	"fmt"
)

type ListQueue struct {
	data *list.List
}

func NewListQueue() *ListQueue {
	return &ListQueue{data: list.New()}
}

func (s *ListQueue) Display() string {
	buf := new(bytes.Buffer)
	buf.WriteString("[")
	for e := s.data.Back(); e != nil; e = e.Prev() {
		_, _ = fmt.Fprintf(buf, "%v", e.Value)
		if e.Prev() != nil {
			buf.WriteString("->")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// Enqueue adds value to the back of a queue
// Time Complexity:
// O(1)
func (s *ListQueue) Enqueue(v any) {
	s.data.PushBack(v) // O(1) it has access via pointer to the back and just add element and swap some pointers
}

// Dequeue removes value from the beginning of a queue
// Time Complexity:
// O(1)
func (s *ListQueue) Dequeue() any {
	elem := s.data.Front() // O(1) direct access to the front via pointer
	if elem == nil {
		return nil
	}

	s.data.Remove(elem) // O(1) as it has direct access to the back of the queue and just swap some pointers

	return elem.Value
}
