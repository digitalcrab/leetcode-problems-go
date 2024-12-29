package main

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// assume that the first element is the majority one
	candidate := nums[0]
	// and give it one vote
	votes := 1

	// check the rest starting from the second
	for _, n := range nums[1:] {
		// let the candidate vote for itself
		if n == candidate {
			votes++
		} else {
			// or against if it's someone else
			votes--
		}

		// votes drop too much, probably we have another candidate coming
		if votes == 0 {
			candidate = n
			votes = 1
		}
	}

	// Explanation [2, 2, 1, 1, 1, 2, 2]
	// init [2]:       candidate = 2, votes = 1
	// step 1 [2]: --> candidate = 2, votes = 2
	// step 2 [1]: --> candidate = 2, votes = 1
	// step 3 [1]: --> candidate = 1, votes = 1
	// step 4 [1]: --> candidate = 1, votes = 2
	// step 5 [2]: --> candidate = 1, votes = 1
	// step 6 [2]: --> candidate = 2, votes = 1

	return candidate
}
