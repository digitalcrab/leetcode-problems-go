package problems

// https://leetcode.com/problems/contiguous-array/
// Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
//
// Example 1:
// Input: [0,1]
// Output: 2
// Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
//
// Example 2:
// Input: [0,1,0]
// Output: 2
// Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
//
// Note: The length of the given binary array will not exceed 50,000.
func zero2minus(v int) int {
	if v == 1 {
		return 1
	}
	return -1
}

func findMaxLengthOfContiguous(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// variables to store result values
	maxsize := -1

	// Create an auxiliary array sunmleft[].
	// sumleft[i] will be sum of array
	// elements from arr[0] to arr[i]
	sumleft := make([]int, len(nums))

	var min, max int

	// Fill sumleft array and get min and max
	// values in it.  Consider 0 values in arr[]
	// as -1

	sumleft[0], min, max = zero2minus(nums[0]), nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		sumleft[i] = sumleft[i-1] + zero2minus(nums[i])
		if sumleft[i] < min {
			min = sumleft[i]
		}
		if sumleft[i] > max {
			max = sumleft[i]
		}
	}

	hash := make(map[int]int)
	for i := 0; i < max-min+1; i++ {
		hash[i] = -1
	}

	for i := 0; i < len(nums); i++ {
		// Case 1: when the subarray starts from
		//         index 0

		if sumleft[i] == 0 {
			maxsize = i + 1
		}

		// Case 2: fill hash table value. If already
		//         filled, then use it

		if hash[sumleft[i]-min] == -1 {
			hash[sumleft[i]-min] = i
		} else {
			if (i - hash[sumleft[i]-min]) > maxsize {
				maxsize = i - hash[sumleft[i]-min]
			}
		}
	}

	if maxsize == -1 {
		return 0
	}
	return maxsize
}
