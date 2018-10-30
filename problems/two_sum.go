package problems

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Time complexity : O(n)
// Space complexity : O(n)
//
func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, n := range nums {
		// check if we have it in the map already, if so then just take an index of it (value of map)
		if pos, ok := hashMap[n]; ok {
			return []int{pos, i}
		}

		// calculate desired missing value and store it in the map to future use
		desired := target - n
		hashMap[desired] = i
	}

	panic("unreachable")
}
