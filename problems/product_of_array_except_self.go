package problems

// https://leetcode.com/problems/product-of-array-except-self/
//
// Given an array nums of n integers where n > 1,
// return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
//
// Example:
// Input:  [1,2,3,4]
// Output: [24,12,8,6]
// Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array
// (including the whole array) fits in a 32 bit integer.
//
// Note: Please solve it without division and in O(n).
//
// Follow up:
// Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)

func productExceptSelf(nums []int) []int {
	// calculate everything from the left of the i-th element
	output := make([]int, len(nums))
	output[0] = 1 // nothing from the left, so 1
	// given example would become [1,1,2,6]
	for i := 1; i < len(nums); i++ {
		output[i] = nums[i-1] * output[i-1]
	}

	// same goes go the right
	right := 1
	// given example would become [24,12,4,1]
	for i := len(nums) - 2; i >= 0; i-- {
		newRight := nums[i+1] * right
		output[i] = output[i] * newRight
		right = newRight
	}

	return output
}
