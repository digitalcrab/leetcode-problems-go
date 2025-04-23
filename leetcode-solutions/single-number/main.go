package main

import "fmt"

// Time Complexity: O(n) where n is number of elements in the array
// Space Complexity: O(1) as we do not allocate any new space
func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		// the usage of XOR is super smart here, just a reminder how does XOR works
		//      A | B | R
		//      0 | 0 | 0
		//      0 | 1 | 1
		// 		1 | 0 | 1
		// 		1 | 1 | 0
		// This means if we XOR identical numbers, lets say 5 (101) and 5 (101) they both result in 0
		// (similar bytes produce zero), if we add some non-duplicate number 3 (011), those bytes
		// gonna stay as they are at the end, that is going to be our result.
		res ^= n
	}
	return res
}

func main() {
	single := singleNumber([]int{2, 2, 1})
	fmt.Printf("Single number is : %d\n", single)
}
