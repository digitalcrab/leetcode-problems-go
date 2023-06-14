package structures

import (
	"testing"
)

func TestStack(t *testing.T) {
	st := NewStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)

	if v := st.Pop(); v != 3 {
		t.Errorf("unexpected result %d -> %d", v, 3)
	}
	if v := st.Pop(); v != 2 {
		t.Errorf("unexpected result %d -> %d", v, 2)
	}
	if v := st.Pop(); v != 1 {
		t.Errorf("unexpected result %d -> %d", v, 1)
	}
	if v := st.Pop(); v != nil {
		t.Errorf("unexpected result %d -> %v", v, nil)
	}
}
