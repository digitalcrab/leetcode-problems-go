package problems

// RemoveElement27 solves https://leetcode.com/problems/remove-element/
//
// Given an integer array nums and an integer val, remove all occurrences
// of val in nums in-place. The order of the elements may be changed.
// Then return the number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal to val
// be k, to get accepted, you need to do the following things:
//
// Change the array nums such that the first k elements of nums contain
// the elements which are not equal to val. The remaining elements of nums
// are not important as well as the size of nums.
// Return k.
//
// Constraints:
//
// 0 <= nums.length <= 100
// 0 <= nums[i] <= 50
// 0 <= val <= 100
//
// Space complexity: O(1) - no extra space is used
// Time complexity: O(n) - we iterate over the array once
func RemoveElement27(nums []int, val int) int {
	idx := 0
	// iterate over the array and move all elements which are
	// not equal to val to the idx position
	// idx will be the number of elements which were moved
	// update idx only when we move an element
	for _, n := range nums {
		if n != val {
			nums[idx] = n
			idx++
		}
	}
	return idx
}
