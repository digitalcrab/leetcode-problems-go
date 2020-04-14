package problems

import "math"

// You are given a string s containing lowercase English letters, and a matrix shift,
// where shift[i] = [direction, amount]:
//   - direction can be 0 (for left shift) or 1 (for right shift).
//   - amount is the amount by which string s is to be shifted.
//   - A left shift by 1 means remove the first character of s and append it to the end.
//   - Similarly, a right shift by 1 means remove the last character of s and add it to the beginning.
// Return the final string after all operations.
//
// Example 1:
// Input: s = "abc", shift = [[0,1],[1,2]]
// Output: "cab"
// Explanation:
//   [0,1] means shift to left by 1. "abc" -> "bca"
//   [1,2] means shift to right by 2. "bca" -> "cab"
//
// Example 2:
// Input: s = "abcdefg", shift = [[1,1],[1,1],[0,2],[1,3]]
// Output: "efgabcd"
// Explanation:
//   [1,1] means shift to right by 1. "abcdefg" -> "gabcdef"
//   [1,1] means shift to right by 1. "gabcdef" -> "fgabcde"
//   [0,2] means shift to left by 2. "fgabcde" -> "abcdefg"
//   [1,3] means shift to right by 3. "abcdefg" -> "efgabcd"
//
// Constraints:
//   1 <= s.length <= 100
//   s only contains lower case English letters.
//   1 <= shift.length <= 100
//   shift[i].length == 2
//   0 <= shift[i][0] <= 1
//   0 <= shift[i][1] <= 100

func stringShift(s string, shift [][]int) string {
	// looks like left shift is cancelled by the right shift
	// this means that we calculate number of shifts and do it only once!
	var totalShift int
	for _, sh := range shift {
		if sh[0] == 1 {
			totalShift += sh[1]
		} else {
			totalShift -= sh[1]
		}
	}

	// if total shift times longer the number of items in the string
	// we shift only leftover, no sense to shift around
	if int(math.Abs(float64(totalShift))) >= len(s) {
		totalShift = totalShift % len(s)
	}

	// nothing to do
	if totalShift == 0 {
		return s
	}

	var shiftedPart, newString string

	// left
	if totalShift < 0 {
		totalShift *= -1
		// shifted part everything from the left up to the given number
		shiftedPart, newString = s[:totalShift], s[totalShift:]
		newString += shiftedPart
	} else { // right
		idx := len(s) - totalShift
		shiftedPart, newString = s[idx:], s[:idx]
		newString = shiftedPart + newString
	}

	return newString
}
