package structures

import "testing"

func TestIntMinHeap(t *testing.T) {
	heap := &intMinHeap{}
	heap.add(10)
	heap.add(8)
	heap.add(6)
	heap.add(4)
	heap.add(2)
	heap.add(-2)

	if v := heap.poll(); v != -2 {
		t.Errorf("unexpected result %d -> %d", v, -2)
	}
	if v := heap.poll(); v != 2 {
		t.Errorf("unexpected result %d -> %d", v, 2)
	}
	heap.poll() // 4
	heap.poll() // 6
	heap.poll() // 8
	if v := heap.poll(); v != 10 {
		t.Errorf("unexpected result %d -> %d", v, 10)
	}
}
