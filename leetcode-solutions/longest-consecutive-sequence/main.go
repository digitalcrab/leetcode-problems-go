package main

import "fmt"

func longestConsecutive(nums []int) int {
	var longest int

	// convert that into the map, also removes duplicates
	sets := map[int]bool{}
	for _, n := range nums { // Time: O(n)
		sets[n] = true
	}

	// loop over map
	for n := range sets { // Time: O(n)
		// check if there is another element smaller than current,
		// if not then we do not need this one as it's not the beginning
		if sets[n-1] {
			continue
		}

		// start looking for all in that sequence
		next := n + 1
		for ; sets[next]; next++ {
			// Time: O(m), where m is the length of sequence, in worst case it could be all numbers again
			// then outer loop can not trigger that loop as often ;)
		}

		// how long was that?
		longest = max(longest, next-n)
	}

	return longest
}

func longestConsecutiveForest(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	forest, sets := newForest(nums)

	// perform union of the siblings
	for n := range sets {
		if _, found := sets[n-1]; found {
			forest.union(n, n-1)
		}
		if _, found := sets[n+1]; found {
			forest.union(n, n+1)
		}
	}

	return forest.maxSize
}

type unionFind struct {
	parent  map[int]int
	size    map[int]int
	maxSize int
}

func newForest(nums []int) (*unionFind, map[int]bool) {
	// start with allocating some space
	sets := make(map[int]bool, len(nums))

	uf := &unionFind{
		parent:  make(map[int]int, len(nums)),
		size:    make(map[int]int, len(nums)),
		maxSize: 1,
	}

	for _, n := range nums { // Time: O(n)
		uf.parent[n] = n
		uf.size[n] = 1 // size is one initially
		sets[n] = true
	}

	return uf, sets
}

// find searches for the parent of the item, and preforms path compression
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	// search for both parents
	rootX := uf.find(x)
	rootY := uf.find(y)

	// if they have same parent then they are in the same set, nothing to do
	if rootX == rootY {
		return
	}

	// make sure that we attach smaller tree to the largest (assume x should be smallest)
	if uf.size[rootX] > uf.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]

	// update max size
	uf.maxSize = max(uf.maxSize, uf.size[rootX])
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	length := longestConsecutive(nums)
	fmt.Printf("Longest consecutive sequence: %d\n", length)
	lengthWithUnion := longestConsecutiveForest(nums)
	fmt.Printf("Longest consecutive sequence: %d\n", lengthWithUnion)
}
