package problems

import (
	"container/heap"
)

// https://leetcode.com/problems/last-stone-weight/
//
// We have a collection of stones, each stone has a positive integer weight.
// Each turn, we choose the two heaviest stones and smash them together.
// Suppose the stones have weights x and y with x <= y.
//
// The result of this smash is:
//   If x == y, both stones are totally destroyed;
//   If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
//
// At the end, there is at most 1 stone left.
// Return the weight of this stone (or 0 if there are no stones left.)
//
// Example 1:
// Input: [2,7,4,1,8,1]
// Output: 1
// Explanation:
//   We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
//   we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
//   we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
//   we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.
//
// Note:
//  1 <= stones.length <= 30
//  1 <= stones[i] <= 1000

type stonesHeap []int

func (h stonesHeap) Len() int {
	return len(h)
}

func (h stonesHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h stonesHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *stonesHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *stonesHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// Time complexity : O(N log N) - Main look multiplied my heap op
// Space complexity : O(N) - a new slice
func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	if len(stones) == 1 {
		return stones[0]
	}

	// place everything into the max heap
	maxHeap := make(stonesHeap, len(stones))
	for i, v := range stones { // Time complexity - O(N)
		maxHeap[i] = v
	}

	heap.Init(&maxHeap) // Time complexity - O(N),

	for maxHeap.Len() > 1 { // Time complexity - O (N - 1) , this loop potentially makes 3 calls to the heap
		// x <= y
		y, x := heap.Pop(&maxHeap).(int), heap.Pop(&maxHeap).(int) // Time complexity - O(log N) + O(log N)

		// If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
		if x != y {
			heap.Push(&maxHeap, y-x) // Time complexity - O(log N)
		}

		// If x == y, both stones are totally destroyed;
		// so pop does it for us
	}

	if maxHeap.Len() == 1 {
		return heap.Pop(&maxHeap).(int)
	}

	return 0
}
