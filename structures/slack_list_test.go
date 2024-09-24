package structures

import "testing"

func TestListStack_Push(t *testing.T) {
	ss := NewListStack()
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)

	if got := ss.Display(); got != "[3, 2, 1]" {
		t.Errorf("Expected stack is different, got: %s", got)
	}
}

func TestListStack_Pop_Peek(t *testing.T) {
	ss := NewListStack()
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)

	if got := ss.Pop(); got != 3 {
		t.Errorf("The recent element expected to be 3 but got %d", got)
	}
	if got := ss.Pop(); got != 2 {
		t.Errorf("The recent element expected to be 2 but got %d", got)
	}
	if got := ss.Pop(); got != 1 {
		t.Errorf("The recent element expected to be 1 but got %d", got)
	}

	ss.Push(42)

	if got := ss.Peek(); got != 42 {
		t.Errorf("The recent element expected to be 42 but got %d", got)
	}
}
