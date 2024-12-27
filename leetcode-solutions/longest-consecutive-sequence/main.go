package main

import "fmt"

func longestConsecutive(nums []int) int {
	// store number to length
	numToLen := make(map[int]int)
	longest := 0

	for _, num := range nums { // Time: O(n) as we go through each element
		// that's we have calculate already so we skip
		if numToLen[num] != 0 {
			continue
		}

		// find known left and right lengths
		left := numToLen[num-1]
		right := numToLen[num+1]

		// calc total length
		sum := left + right + 1

		// update total for cur rent number
		numToLen[num] = sum

		// update to the left most and right most items
		numToLen[num-left] = sum
		numToLen[num+right] = sum

		// update longest
		longest = max(longest, sum)
	}

	return longest
}

func longestConsecutiveHash(nums []int) int {
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
	sets := make(map[int]bool)

	uf := &unionFind{
		parent:  make(map[int]int),
		size:    make(map[int]int),
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
	length := longestConsecutiveHash(nums)
	fmt.Printf("Longest consecutive sequence: %d\n", length)
	lengthWithUnion := longestConsecutiveForest(nums)
	fmt.Printf("Longest consecutive sequence: %d\n", lengthWithUnion)
}
