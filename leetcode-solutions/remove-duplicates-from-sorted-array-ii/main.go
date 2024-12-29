package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// remember the first element
	seen := nums[0]
	// and set the current index to the second element (starts from 0)
	current := 1
	// also remember how may of those we have
	inserted := 1

	// start looping from the second
	for _, n := range nums[1:] {
		// handle the case of compliantly new value
		if n != seen {
			nums[current] = n
			current++
			seen = n
			inserted = 1
		} else if inserted < 2 {
			// the same value but we've added only 1 so far
			nums[current] = n
			current++
			inserted++
		}
	}

	return current
}
