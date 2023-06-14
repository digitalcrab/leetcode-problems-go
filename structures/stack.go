package structures

type Stack struct {
	list *LinkedList
}

func NewStack() *Stack {
	return &Stack{list: NewLinkedList()}
}

// Push adds an element to the collection, and
func (s *Stack) Push(v any) {
	s.list.InsertBeginning(v)
}

// Pop removes the most recently added element that was not yet removed.
func (s *Stack) Pop() (v any) {
	return s.list.RemoveBeginning()
}
