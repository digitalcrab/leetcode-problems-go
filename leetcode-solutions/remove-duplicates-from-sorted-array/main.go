package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// remember the first element
	seen := nums[0]
	// and set the current index to the second element (starts from 0)
	current := 1

	// start looping from the second
	for _, n := range nums[1:] { // Time: O(n)
		// if elem is not the same as we know
		if n != seen {
			// move it to the next position
			nums[current] = n
			// inc the position counter
			current++
			// change the last seen
			seen = n
		}
	}

	// return how many we have
	return current
}
