package problems

// RotateArray189Reverse solves https://leetcode.com/problems/rotate-array/
//
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
// Constraints:
// 1 <= nums.length <= 10^5
// -2^31 <= nums[i] <= 2^31 - 1
// 0 <= k <= 10^5
//
// Space complexity: O(1)
// Time complexity: O(n+k+(n-k)) = O(n)
func RotateArray189Reverse(nums []int, k int) {
	// reverse elements in the array starting from index `from` to index `to`
	reverse := func(from, to int) {
		for from < to {
			nums[from], nums[to] = nums[to], nums[from]
			from++
			to--
		}
	}

	// if k is greater than the length of the array, then we need to rotate only k % len(nums) elements
	if k > len(nums) {
		k = k % len(nums)
	}

	// reverse all elements in the array
	reverse(0, len(nums)-1)

	// reverse first k elements
	reverse(0, k-1)

	// reverse last n-k elements
	reverse(k, len(nums)-1)
}

// RotateArray189ShiftingToRight same as RotateArray189. No extra memory used, elements are shifted to the right.
// Space complexity: O(1)
// Time complexity: O(n*k)
func RotateArray189ShiftingToRight(nums []int, k int) {
	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		// shift right all elements
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

// RotateArray189ExtraMemory same as RotateArray189 but uses extra memory
// Space complexity: O(n)
// Time complexity: O(n)
func RotateArray189ExtraMemory(nums []int, k int) {
	// copy last k elements to the beginning of the array
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}
