package structures

import "testing"

func TestIntMaxHeap(t *testing.T) {
	heap := &intMaxHeap{}
	heap.add(10)
	heap.add(8)
	heap.add(6)
	heap.add(4)
	heap.add(2)
	heap.add(-2)

	if v := heap.poll(); v != 10 {
		t.Errorf("unexpected result %d -> %d", v, 10)
	}
	if v := heap.poll(); v != 8 {
		t.Errorf("unexpected result %d -> %d", v, 8)
	}
	heap.poll() // 6
	heap.poll() // 4
	heap.poll() // 2
	if v := heap.poll(); v != -2 {
		t.Errorf("unexpected result %d -> %d", v, -2)
	}
}
