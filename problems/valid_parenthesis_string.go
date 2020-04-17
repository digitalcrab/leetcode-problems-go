package problems

import "math"

// https://leetcode.com/problems/valid-parenthesis-string/solution/
//
// Given a string containing only three types of characters: '(', ')' and '*',
// write a function to check whether this string is valid.
//
// We define the validity of a string by these rules:
//  - Any left parenthesis '(' must have a corresponding right parenthesis ')'.
//  - Any right parenthesis ')' must have a corresponding left parenthesis '('.
//  - Left parenthesis '(' must go before the corresponding right parenthesis ')'.
//  - '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
//  - An empty string is also valid.
//
// Example 1:
// Input: "()"
// Output: True
//
// Example 2:
// Input: "(*)"
// Output: True
//
// Example 3:
// Input: "(*))"
// Output: True
//
// Note:
//   The string size will be in the range [1, 100].

func checkValidString(s string) bool {
	// empty is a correct string, as well as '*'
	if s == "" || s == "*" {
		return true
	}
	// then every char should have a pair, so at least 2 should be present
	if len(s) < 2 {
		return false
	}

	// we keep track of smallest and largest possible number of open left brackets '('
	var smallest, largest int

	for _, c := range s {
		// we definitely have at least one
		if c == '(' {
			smallest += 1
		} else {
			// otherwise we assume it's ')' (not really counting on '*' because it's
			// a smallest possible number)
			smallest -= 1
		}

		// in the same time

		// if it's not a ')' so it could be '(' or '*'
		// largest possible number should grow
		// (same assumption as pref if statement but wider scope)
		if c != ')' {
			largest += 1
		} else {
			// this could be ')' or '*'
			largest -= 1
		}

		// current prefix can't be made valid no matter what our choices are
		if largest < 0 {
			break
		}

		// can never have less than 0 open left brackets
		smallest = int(math.Max(float64(smallest), 0))
	}

	// we check ho much left
	return smallest == 0
}
