package problems

// https://leetcode.com/problems/backspace-string-compare/
//
// Given two strings S and T, return if they are equal when both are typed into empty text editors.
// # means a backspace character.
//
// Example 1:
// Input: S = "ab#c", T = "ad#c"
// Output: true
// Explanation: Both S and T become "ac".
//
// Example 2:
// Input: S = "ab##", T = "c#d#"
// Output: true
// Explanation: Both S and T become "".
//
// Example 3:
// Input: S = "a##c", T = "#a#c"
// Output: true
// Explanation: Both S and T become "c".
//
// Example 4:
// Input: S = "a#c", T = "b"
// Output: false
// Explanation: S becomes "c" while T becomes "b".
//
// Note:
//   1 <= S.length <= 200
//   1 <= T.length <= 200
//   S and T only contain lowercase letters and '#' characters.
//
// Follow up:
//   Can you solve it in O(N) time and O(1) space?

func backspaceCompare(S string, T string) bool {
	return backspaceCompareProcess(S) == backspaceCompareProcess(T)
}

func backspaceCompareProcess(input string) string {
	if len(input) == 0 {
		return ""
	}

	output := make([]rune, 0, len(input))

	for _, ch := range input {
		if ch == '#' {
			if len(output) > 0 {
				output = output[:len(output)-1]
			}
		} else {
			output = append(output, ch)
		}
	}

	return string(output)
}
