package problems

// MergeSortedArray88 solves https://leetcode.com/problems/merge-sorted-array/
//
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// and two integers m and n, representing the number of elements in nums1 and nums2
// respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead
// be stored inside the array nums1. To accommodate this, nums1 has a
// length of m + n, where the first m elements denote the elements that should
// be merged, and the last n elements are set to 0 and should be ignored.
// nums2 has a length of n.
//
// Constraints:
//
//	nums1.length == m + n
//	nums2.length == n
//	0 <= m, n <= 200
//	1 <= m + n <= 200
//	-10^9 <= nums1[i], nums2[j] <= 10^9
//
// Space complexity: O(1) - no extra space is used
// Time complexity: O(n + m) - we iterate over both arrays once, from the end to the beginning
func MergeSortedArray88(nums1 []int, m int, nums2 []int, n int) {
	// base cases if nums1 is empty
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	// base cases if nums2 is empty
	if n == 0 {
		return
	}

	// start from the end of the arrays
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			// nums1[m-1] is bigger than nums2[n-1], it has to go to the end of nums1,
			// and we move m one step to the left
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
