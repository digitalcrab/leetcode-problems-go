package main

// Time: O(n) + O(n) = O(n)
func rotate(nums []int, k int) {
	// if we loop over the len, then we do only like the rest ;)
	if k > len(nums) {
		k = k % len(nums)
	}

	// nums = [1,2,3,4,5,6,7], k = 3
	// expected = [5,6,7,1,2,3,4]

	// the idea is to rotate the whole slice
	// nums = [7,6,5,4,3,2,1]
	// this way last k elems are already in front of the slice
	reverse(nums, 0, len(nums)-1) // Time: O(n)

	// now rotate them
	// nums = [5,6,7,4,3,2,1]
	reverse(nums, 0, k-1) // Time: O(k)

	// then it looks exactly as we need
	// fix the rest by rotation
	// nums = [5,6,7,1,2,3,4]
	reverse(nums, k, len(nums)-1) // Time: O(n-k)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// Time: O(n)
// Space: O(n)
func rotateV2(nums []int, k int) {
	// if we loop over the len, then we do only like the rest ;)
	if k > len(nums) {
		k = k % len(nums)
	}

	var stack []int

	// add last k elements first
	for i := len(nums) - k; i < len(nums); i++ {
		stack = append(stack, nums[i])
	}

	// add the rest of them
	for i := 0; i < len(nums)-k; i++ {
		stack = append(stack, nums[i])
	}

	copy(nums, stack)
}
