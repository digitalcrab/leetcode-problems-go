package problems

import "math"

// Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
//
// Example 1:
// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
//
// Example 2:
// Input: "cbbd"
// Output: "bb"
//
// Hint: A palindrome is a word, number, phrase, or other sequence of characters which reads the same
// backward as forward, such as madam, racecar.

// Time complexity : O(n)
// Space complexity : O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	// initial boundaries for the longest palindrome
	startIdx, endIdx := 0, 0

	for i := 0; i < len(s); i++ { // O(n)
		// handle the case like "racecar" where we start from the middle char "e"
		lenOne := float64(expandFromTheMiddle(s, i, i)) // O(n)
		// "normal" case for example "abba" where we start from both "b"
		lenTwo := float64(expandFromTheMiddle(s, i, i+1)) // O(n)

		ln := int(math.Max(lenOne, lenTwo))

		// check if the one we found is bigger that the one known before
		if ln > endIdx-startIdx {
			// update boundaries
			// example "babad": i = 1, len = 3
			// startIdx = 1 - (3-1)/2 = 0
			// endIdx = 1 + 3/2 = 2
			startIdx = i - (ln-1)/2
			endIdx = i + ln/2
		}
	}

	// example "babad": startIdx = 0, endIdx = 2
	// s[startIdx : endIdx+1] = "bab", go works like [n:m)
	return s[startIdx : endIdx+1]
}

// O(n)
func expandFromTheMiddle(s string, leftIdx, rightIdx int) int {
	if len(s) == 0 || leftIdx > rightIdx {
		return 0
	}

	// going from middle point to the left and right simultaneously
	// and if char is the same keep moving indexes to the left and right
	for leftIdx >= 0 && rightIdx < len(s) && s[leftIdx] == s[rightIdx] {
		leftIdx--
		rightIdx++
	}

	// return the length of the palindrome,
	// example "babad": rightIdx = 3, leftIdx = -1
	// 3 - (-1) - 1 = 3

	return rightIdx - leftIdx - 1
}
