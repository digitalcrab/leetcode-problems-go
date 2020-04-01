package problems

// https://leetcode.com/problems/single-number/
// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
//
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
//
// Example 1:
// Input: [2,2,1]
// Output: 1

// Example 2:
// Input: [4,1,2,1,2]
// Output: 4

func singleNumber(nums []int) int {
	return singleNumberXor(nums)
}

// Runtime complexity: O(n)
// Memory: O(n)
//
// A bitwise XOR is a binary operation that takes two bit patterns of equal length and performs
// the logical exclusive OR operation on each pair of corresponding bits.
// The result in each position is 1 if only the first bit is 1 or only the second bit is 1,
// but will be 0 if both are 0 or both are 1.
//
// In this we perform the comparison of two bits, being 1 if the two bits are different, and 0 if they are the same.
//
// 4 - 00000100
// 1 - 00000001 __ 00000101
// 2 - 00000010 __ 00000111
// 1 - 00000001 __ 00000110
// 2 - 00000010 __ 00000100 = 4
//
func singleNumberXor(nums []int) int {
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}

// Runtime complexity: O(n)
// Memory: O(n2)
func singleNumberHash(nums []int) int {
	// extra memory comes from here
	seen := make(map[int]int)

	for _, n := range nums {
		if _, ok := seen[n]; !ok {
			seen[n] = 1
		} else {
			seen[n]++
		}
	}

	for e, c := range seen {
		if c == 1 {
			return e
		}
	}

	panic("not reachable")
}
