package problems

import "testing"

func TestMinStack(t *testing.T) {
	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)

	if v := s.GetMin(); v != -3 {
		t.Errorf("unexpected resutl %d, wants %d", v, -3)
	}

	s.Pop()

	if v := s.Top(); v != 0 {
		t.Errorf("unexpected resutl %d, wants %d", v, 0)
	}
	if v := s.GetMin(); v != -2 {
		t.Errorf("unexpected resutl %d, wants %d", v, -2)
	}
}
