package structures

import (
	"testing"
)

func TestSlicesStack(t *testing.T) {
	ss := make(SlicesStack, 0)
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
