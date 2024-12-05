package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// create map to store unique elements, struct is not going to take any space
	seen := make(map[int]struct{})
	for _, n := range nums {
		// if we've seen any of those numbers, then no need to check more
		if _, exists := seen[n]; exists {
			return true
		}
		// save the fact that we've seen
		seen[n] = struct{}{}
	}
	return false
}

func main() {
	fmt.Printf("[1,2,3,1] = %t\n", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Printf("[1,2,3,4] = %t\n", containsDuplicate([]int{1, 2, 3, 4}))
}
