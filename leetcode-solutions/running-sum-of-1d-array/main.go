package main

import "fmt"

// Time Complexity: O(n) where n is a number of elements in the nums array
// Space Complexity: O(1) as we reuse the array
func runningSum(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// the running sum of a first element is the first element inself
	// start looping from the second and add it to the previous
	// prev is a running sum that we have calculated already
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := runningSum(nums)
	fmt.Printf("Running sum: %v\n", result)
}
