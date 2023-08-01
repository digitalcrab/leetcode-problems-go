package problems

// MajorityElement169 solves https://leetcode.com/problems/majority-element/
//
// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
//
// Constraints:
//
// n == nums.length
// 1 <= n <= 5 * 10^4
// -10^9 <= nums[i] <= 10^9
//
// Follow-up: Could you solve the problem in linear time and in O(1) space?
func MajorityElement169(nums []int) int {
	return majorityElement169BoyerMooreMajorityVotingAlgorithm(nums)
}

// Space complexity: O(n)
// Time complexity: O(n)
func majorityElement169Naive(nums []int) int {
	majority := len(nums) / 2
	appears := make(map[int]int)
	for _, num := range nums {
		appears[num]++
		if appears[num] > majority {
			return num
		}
	}
	return -1
}

// The Boyer-Moore voting algorithm is one of the popular optimal algorithms
// which is used to find the majority element among the given elements that have
// more than N/ 2 occurrences.
//
// @see https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm/
//
// Space complexity: O(1)
// Time complexity: O(n)
func majorityElement169BoyerMooreMajorityVotingAlgorithm(nums []int) int {
	candidate := nums[0]
	votes := 1
	for _, n := range nums[1:] {
		if votes == 0 {
			candidate = n
			votes = 1
		} else {
			if candidate == n {
				votes++
			} else {
				votes--
			}
		}
	}
	return candidate
}
