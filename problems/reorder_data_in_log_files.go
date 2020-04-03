package problems

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/reorder-data-in-log-files/
//
// You have an array of logs.  Each log is a space delimited string of words.
//
// For each log, the first word in each log is an alphanumeric identifier.  Then, either:
//  - Each word after the identifier will consist only of lowercase letters, or;
//  - Each word after the identifier will consist only of digits.
//
// We will call these two varieties of logs letter-logs and digit-logs.
// It is guaranteed that each log has at least one word after its identifier.
//
// Reorder the logs so that all of the letter-logs come before any digit-log.
// The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties.
// The digit-logs should be put in their original order.
//
// Return the final order of the logs.
//
// Example 1:
// Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
// Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
//
// Constraints:
// 0 <= logs.length <= 100
// 3 <= logs[i].length <= 100
// logs[i] is guaranteed to have an identifier, and a word after the identifier.

type byReorderDataInLogFile []string

func (s byReorderDataInLogFile) Len() int {
	return len(s)
}

func isDigit(r byte) bool {
	return '0' <= r && r <= '9'
}

func (s byReorderDataInLogFile) Less(i, j int) bool {
	s1 := strings.SplitN(s[i], " ", 2)
	s2 := strings.SplitN(s[j], " ", 2)

	isDigit1 := isDigit(s1[1][0])
	isDigit2 := isDigit(s2[1][0])

	if !isDigit1 && !isDigit2 {
		if s1[1] == s2[1] {
			return s[i] < s[j]
		}
		return s1[1] < s2[1]
	}

	if isDigit1 && isDigit2 {
		return false
	}
	if isDigit1 && !isDigit2 {
		return false
	}

	return true
}

func (s byReorderDataInLogFile) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func reorderLogFiles(logs []string) []string {
	res := make([]string, len(logs))
	copy(res, logs)
	sort.Sort(byReorderDataInLogFile(res)) // O(n*log(n))
	return res
}
