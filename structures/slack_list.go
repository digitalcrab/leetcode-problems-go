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
func (s *ListStack) Push(v any) {
	s.data.PushFront(v)
}

// Pop removes the most recent added element
func (s *ListStack) Pop() any {
	elem := s.data.Front()
	if elem == nil {
		return nil
	}
	s.data.Remove(elem)
	return elem.Value
}

// Peek returns the recent element without modifying the stack
func (s *ListStack) Peek() any {
	elem := s.data.Front()
	if elem == nil {
		return nil
	}
	return elem.Value
}
