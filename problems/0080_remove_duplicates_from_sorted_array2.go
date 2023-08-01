package problems

// RemoveDuplicatesFromSortedArray80 solves https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
//
// Given an integer array nums sorted in non-decreasing order, remove some
// duplicates in-place such that each unique element appears at most twice.
// The relative order of the elements should be kept the same.
//
// Since it is impossible to change the length of the array in some languages,
// you must instead have the result be placed in the first part of the array nums.
// More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final result.
// It does not matter what you leave beyond the first k elements.
//
// Return k after placing the final result in the first k slots of nums.
//
// Do not allocate extra space for another array.
// You must do this by modifying the input array in-place with O(1) extra memory.
//
// Constraints:
//
// 1 <= nums.length <= 3 * 10^4
// -10^4 <= nums[i] <= 10^4
// nums is sorted in non-decreasing order.
//
// Space complexity: O(1) - no extra space is used
// Time complexity: O(n) - we iterate over the array once
func RemoveDuplicatesFromSortedArray80(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// define first element as the last one we have seen
	last := nums[0]
	idx := 1
	inserted := 1

	// start from the second element
	for _, n := range nums[1:] {
		// the element is the same as the last one and we have inserted it less than 2 times
		if n == last && inserted < 2 {
			nums[idx] = n
			idx++
			inserted++
		}

		// this is a new element, we need to insert it and reset the counter
		if n != last {
			nums[idx] = n
			idx++
			inserted = 1
		}

		last = n
	}
	return idx
}
