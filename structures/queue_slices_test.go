package structures

import (
	"testing"
)

func TestSlicesQueue_Enqueue(t *testing.T) {
	q := new(SlicesQueue)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if got := q.Display(); got != "[3->2->1]" {
		t.Errorf("Expected queue is different, got: %s", got)
	}
}

func TestSlicesQueue_Dequeue(t *testing.T) {
	q := new(SlicesQueue)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if got := q.Dequeue(); got != 1 {
		t.Errorf("The next element expected to be 1 but got %d", got)
	}
	if got := q.Dequeue(); got != 2 {
		t.Errorf("The next element expected to be 2 but got %d", got)
	}
	if got := q.Dequeue(); got != 3 {
		t.Errorf("The next element expected to be 3 but got %d", got)
	}

	q.Enqueue(42)

	if got := q.Dequeue(); got != 42 {
		t.Errorf("The next element expected to be 42 but got %d", got)
	}
}
