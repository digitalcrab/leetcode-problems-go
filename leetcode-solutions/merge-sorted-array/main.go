package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// first cover the edge case when m expected to be empty
	if m == 0 {
		// simply copy everything from nums2 up until n elements
		copy(nums1, nums2[:n])
		return
	}

	// second edge case when n is empty
	if n == 0 {
		// nothing to do, just lets make sure we have no more than m elements in the nums1
		nums1 = nums1[:m]
		return
	}

	// we gonna run till we run out of elements in nums2
	for n > 0 {
		// if there are elements in the nums1 and
		// last element of nums1 is bigger than the last of nums2
		if m > 0 && nums1[m-1] > nums2[n-1] {
			// move it to the last position
			nums1[m+n-1] = nums1[m-1]
			// and move the cursor to the next from the end
			m--
		} else {
			// if no elements left in nums1 or nums2 last element is bigger
			// than the last of nums1, move the last of nums2 to the nums1
			nums1[m+n-1] = nums2[n-1]
			// move cursor
			n--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Printf("Nums1 after merge: %v\n", nums1)
}
