package structures

// HeapArray is a tree-based data structure
// Example: [ 0, 1, 2, 3, 4, 5, 6 ]
// Where:
//
//	     0
//	    / \
//	  1     2
//	 / \   / \
//	3   4 5   6
//
// for each node the left one is `idx*2`+1 and the right `idx*2+2`
// parent could be found like this:  `(idx-1)/2`
type HeapArray struct {
	data []any
	less LessFunc // this func makes the deposition what is smaller
}

func NewHeapArray(less LessFunc) *HeapArray {
	h := &HeapArray{
		data: make([]any, 0),
		less: less,
	}
	return h
}

func (h *HeapArray) Push(v any) {
	h.data = append(h.data, v)
	idx := len(h.data) - 1
	h.bubbleUp(idx)
}

// bubbleUp pushes the `idx` up to find the correct spot
func (h *HeapArray) bubbleUp(idx int) {
	for {
		parent := (idx - 1) / 2
		// if parent is smaller
		if idx == parent || h.less(h.data[parent], h.data[idx]) {
			break
		}
		// swap
		h.data[parent], h.data[idx] = h.data[idx], h.data[parent]
		// move to the next
		idx = parent
	}
}

func (h *HeapArray) Pop() any {
	if len(h.data) == 0 {
		return nil
	}
	// get the top element
	v := h.data[0]
	// swap it with the last
	idx := len(h.data) - 1
	h.data[0], h.data[idx] = h.data[idx], h.data[0]
	// remove the last
	h.data = h.data[:idx]
	// basically the biggest element
	h.bubbleDown(0)
	return v
}

func (h *HeapArray) bubbleDown(idx int) {
	maxIdx := len(h.data) - 1
	for {
		leftIdx := idx*2 + 1
		rightIdx := idx*2 + 2 // or leftIdx + 1

		// check boundaries
		if leftIdx > maxIdx || leftIdx < 0 {
			break
		}

		// assume for now that left is smaller
		idxToSwap := leftIdx

		// check if right is smaller than left (looking for smaller child)
		if rightIdx < maxIdx && h.less(h.data[rightIdx], h.data[leftIdx]) {
			idxToSwap = rightIdx // now we know that we have right and it's smaller than left
		}

		// if idx is already smaller than smallest child, then we are good
		if h.less(h.data[idx], h.data[idxToSwap]) {
			break
		}

		// swap instead
		h.data[idxToSwap], h.data[idx] = h.data[idx], h.data[idxToSwap]

		// lets check the next iteration over the smallest index of what we know
		idx = idxToSwap
	}
}
