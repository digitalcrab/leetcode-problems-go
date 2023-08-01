package problems

// RemoveDuplicatesFromSortedArray26 solves https://leetcode.com/problems/remove-duplicates-from-sorted-array/
//
// Given an integer array nums sorted in non-decreasing order,
// remove the duplicates in-place such that each unique element appears
// only once. The relative order of the elements should be kept the same.
// Then return the number of unique elements in nums.
//
// Consider the number of unique elements of nums to be k, to get accepted,
// you need to do the following things:
//
// Change the array nums such that the first k elements of nums contain
// the unique elements in the order they were present in nums initially.
// The remaining elements of nums are not important as well as the size of nums.
//
// Return k.
//
// Constraints:
//
// 1 <= nums.length <= 3 * 104
// -100 <= nums[i] <= 100
// nums is sorted in non-decreasing order.
//
// Space complexity: O(1) - no extra space is used
// Time complexity: O(n) - we iterate over the array once
func RemoveDuplicatesFromSortedArray26(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// define first element as the last one we have seen
	last := nums[0]
	idx := 1
	// start from the second element
	for _, n := range nums[1:] {
		// move the element to the idx position always
		nums[idx] = n
		// update idx only when last element is not equal to the current one
		if n != last {
			idx++
		}
		last = n
	}
	return idx
}
