package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	// here we're going to store the values that we have visited already
	// where key is a value that we have visited and index is value in the map
	memo := make(map[int]int)

	for idx, current := range nums {
		// calculate how much we are missing from `current` to `target`
		reminder := target - current
		// check if in `memo` we see the `reminder`, this means that we've seen
		// the number in the past and it's exactly what we need to add to `current`
		// to get `target`
		seenIdx, found := memo[reminder]
		if found {
			// we have it in `memo` to return current `idx` and index that we've seen
			return []int{seenIdx, idx}
		}
		// if not found simply remember the current value and an index
		memo[current] = idx
	}

	// should not happen, but still return nil if no solution is found
	return nil
}

func main() {
	// Example usage
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	fmt.Println(result) // Output: [0, 1]
}
