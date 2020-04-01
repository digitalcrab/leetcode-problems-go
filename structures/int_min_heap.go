package structures

type intMinHeap struct {
	data []int
}

// removal of the root node works as a replacement of the root value with the very last
// value and then rearrange the heap in the way so the new root goes down to the proper place
func (h *intMinHeap) poll() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}

	// get the value of the first element
	v := h.data[0]
	// replace first with the very last
	h.swap(0, len(h.data)-1)
	// downsize the array
	h.data = h.data[:len(h.data)-1]
	// rearrange top
	h.heapifyDown()

	return v
}

// add works simple as adding value to the end of the list and move this value up until it finds proper place
func (h *intMinHeap) add(v int) {
	h.data = append(h.data, v)
	h.heapifyUp()
}

func (h *intMinHeap) heapifyDown() {
	idx := 0

	// if there is no left child then certainly no right child
	for h.hasLeft(idx) {
		smallerChildIdx := h.leftIdx(idx)

		// if we have right and it's smaller then left, swap the index
		if h.hasRight(idx) && h.data[h.rightIdx(idx)] < h.data[smallerChildIdx] {
			smallerChildIdx = h.rightIdx(idx)
		}

		// if current is smaller than both of the children, all fine we keep it here
		if h.data[idx] < h.data[smallerChildIdx] {
			break
		}

		// not the case, then we move down the tree
		h.swap(idx, smallerChildIdx)
		idx = smallerChildIdx
	}
}

func (h *intMinHeap) heapifyUp() {
	idx := len(h.data) - 1

	// has parent and parent is bigger than current
	for h.hasParent(idx) && h.data[h.parentIdx(idx)] > h.data[idx] {
		// swap and get parent idx
		h.swap(h.parentIdx(idx), idx)
		idx = h.parentIdx(idx)
	}
}

func (h *intMinHeap) swap(idx1, idx2 int) {
	h.data[idx1], h.data[idx2] = h.data[idx2], h.data[idx1]
}

func (h *intMinHeap) parentIdx(idx int) int {
	return (idx - 1) / 2
}

func (h *intMinHeap) leftIdx(idx int) int {
	return idx*2 + 1
}

func (h *intMinHeap) rightIdx(idx int) int {
	return idx*2 + 2
}

func (h *intMinHeap) hasParent(idx int) bool {
	return h.parentIdx(idx) >= 0
}

func (h *intMinHeap) hasLeft(idx int) bool {
	return h.leftIdx(idx) < len(h.data)
}

func (h *intMinHeap) hasRight(idx int) bool {
	return h.rightIdx(idx) < len(h.data)
}
