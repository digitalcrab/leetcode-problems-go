package structures

import (
	"testing"
)

func TestHeapArray_Min_PushPop(t *testing.T) {
	less := func(a, b any) bool {
		return a.(int) < b.(int)
	}

	// heap with simple less function for ints
	h := NewHeapArray(less)
	h.Push(5)

	if got := h.Pop(); got != 5 {
		t.Errorf("unexpected min value: %v", got)
	}

	h.Push(6)
	h.Push(1)
	h.Push(-1)
	h.Push(18)
	h.Push(5)

	if got := h.Pop(); got != -1 {
		t.Errorf("unexpected min value: %v", got)
	}
	if got := h.Pop(); got != 1 {
		t.Errorf("unexpected min value: %v", got)
	}
	if got := h.Pop(); got != 5 {
		t.Errorf("unexpected min value: %v", got)
	}
	if got := h.Pop(); got != 6 {
		t.Errorf("unexpected min value: %v", got)
	}
	if got := h.Pop(); got != 18 {
		t.Errorf("unexpected min value: %v", got)
	}
}

func TestHeapArray_Max_PushPop(t *testing.T) {
	less := func(a, b any) bool {
		return a.(int) > b.(int)
	}

	// heap with simple less function for ints
	h := NewHeapArray(less)
	h.Push(5)

	if got := h.Pop(); got != 5 {
		t.Errorf("unexpected min value: %v", got)
	}

	h.Push(6)
	h.Push(1)
	h.Push(-1)
	h.Push(18)
	h.Push(5)

	if got := h.Pop(); got != 18 {
		t.Errorf("unexpected min value: %v", got)
	}
	if got := h.Pop(); got != 6 {
		t.Errorf("unexpected min value: %v", got)
	}
	if got := h.Pop(); got != 5 {
		t.Errorf("unexpected min value: %v", got)
	}
	if got := h.Pop(); got != 1 {
		t.Errorf("unexpected min value: %v", got)
	}
	if got := h.Pop(); got != -1 {
		t.Errorf("unexpected min value: %v", got)
	}
}
