package structures

import (
	"bytes"
	"container/list"
	"fmt"
)

type ListStack struct {
	data *list.List
}

func NewListStack() *ListStack {
	return &ListStack{data: list.New()}
}

func (s *ListStack) Display() string {
	buf := new(bytes.Buffer)
	buf.WriteString("[")
	for e := s.data.Front(); e != nil; e = e.Next() {
		_, _ = fmt.Fprintf(buf, "%v", e.Value)
		if e.Next() != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// Push adds an element to the collection
// Time Complexity:
// O(1) simply changes pointers of next and prev
func (s *ListStack) Push(v any) {
	s.data.PushFront(v)
}

// Pop removes the most recent added element
// Time Complexity:
// O(1) involves only simple operations
func (s *ListStack) Pop() any {
	elem := s.data.Front() // O(1) returns top element by the reference
	if elem == nil {
		return nil
	}
	s.data.Remove(elem) // O(1) changing a few pointers (next + prev)
	return elem.Value
}

// Peek returns the recent element without modifying the stack
// Time Complexity:
// O(1) involves only simple operations
func (s *ListStack) Peek() any {
	elem := s.data.Front()
	if elem == nil {
		return nil
	}
	return elem.Value
}
