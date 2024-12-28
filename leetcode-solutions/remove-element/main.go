package main

import "fmt"

// Time Complexity: O(n) because of loop over nums
func removeElement(nums []int, val int) int {
	// the idea is to have cursor that indicates elements not eq to k
	var cursor int

	// while looping over elements
	for _, n := range nums {
		// and see the one we need to keep
		if n != val {
			// move it to the cursor
			nums[cursor] = n
			// and inc it
			cursor++
		}
	}

	// explanation
	// nums = [3,2,2,3], val = 2
	// iteration 1: cursor = 0, nums = [3,2,2,3], 3 goes (stays) to the pos 0, cursor = 1
	// iteration 2: cursor = 1, nums = [3,2,2,3], 2 skip, cursor = 1
	// iteration 2: cursor = 1, nums = [3,2,2,3], 2 skip, cursor = 1
	// iteration 2: cursor = 1, nums = [3,2,2,3], 3 goes to 1, cursor = 2, nums = [3,3,2,3]
	// assurance happens for the first cursor (2) elements

	// then return how many we have moved
	return cursor
}

// Time Complexity: O(n) + O(min(n,m)) = O(n)
func removeElementStack(nums []int, val int) int {
	var stack []int

	for _, n := range nums { // Time: O(n)
		if n != val {
			stack = append(stack, n)
		}
	}

	return copy(nums, stack) // Time: O(min(n,m)) where n number of elements in nums and m number of elements in stack
}

func main() {
	nums1 := []int{3, 2, 2, 3}
	got1 := removeElement(nums1, 3)
	fmt.Printf("Leftovers: %v\n", nums1[:got1])
	nums2 := []int{3, 2, 2, 3}
	got2 := removeElementStack(nums2, 3)
	fmt.Printf("Leftovers: %v\n", nums2[:got2])
}
