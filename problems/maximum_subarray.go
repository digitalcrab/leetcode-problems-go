package problems

import "math"

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
//
// Time complexity : O(n)
// Space complexity : O(1)
//
func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	biggest := math.MinInt64
	current := 0

	for _, n := range nums {
		current = current + n

		if biggest < current {
			biggest = current
		}

		// does not make sense to sum negative numbers because they sum will lead to lower value
		if current < 0 {
			current = 0
		}
	}

	return biggest
}
