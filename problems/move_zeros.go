package problems

// https://leetcode.com/problems/move-zeroes/
//
// Given an array nums, write a function to move all 0's to the end of it while maintaining
// the relative order of the non-zero elements.
//
// Example:
// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Note:
//  - You must do this in-place without making a copy of the array.
//  - Minimize the total number of operations.

// Time Complexity: O(n) - loop over all elements
// Space Complexity : O(1) - only stack local variables
func moveZeroes(nums []int) {
	for i, lastNonZeroIdx := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZeroIdx] = nums[lastNonZeroIdx], nums[i]
			lastNonZeroIdx++
		}
	}
}
