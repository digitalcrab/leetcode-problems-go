package problems

type intStackNode struct {
	prev  *intStackNode
	value int
}

// advantage of this implementation is that you do not have to know the size of the stack to avoid unnecessary allocations
// (if to compare with pure implementation on slices, see `valid_parentheses.go`)
// disadvantage is an overhead in 4 bytes per element
type intStack struct {
	node *intStackNode
}

func (s *intStack) pop() int {
	if node := s.node; node != nil {
		s.node = node.prev
		return node.value
	}
	return 0
}

func (s *intStack) push(v int) {
	s.node = &intStackNode{value: v, prev: s.node}
}
