package structures

type Stack struct {
	list *LinkedList
}

func NewStack() *Stack {
	return &Stack{list: NewLinkedList()}
}

// Push adds an element to the collection, and
func (s *Stack) Push(v any) {
	s.list.Top = &LinkedListNode{
		Val:  v,
		Next: s.list.Top,
	}
}

// Pop removes the most recently added element that was not yet removed.
func (s *Stack) Pop() (v any) {
	if s.list.Top == nil {
		return nil
	}
	v, s.list.Top = s.list.Top.Val, s.list.Top.Next
	return
}
