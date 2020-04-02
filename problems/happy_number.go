package problems

// https://leetcode.com/problems/happy-number/
// Write an algorithm to determine if a number is "happy".
//
// A happy number is a number defined by the following process:
// - Starting with any positive integer
// - Replace the number by the sum of the squares of its digits
// - Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
//
// Those numbers for which this process ends in 1 are happy numbers.
//
// Example 1:
// Input: 19
// Output: true
// Explanation:
//  1^2 + 9^2 = 82
//  8^2 + 2^2 = 68
//  6^2 + 8^2 = 100
//  1^2 + 0^2 + 0^2 = 1
//
// Example 1:
// Input: 18
// Output: false
// Explanation:
//  1^2 + 8^2 = 65
//  6^2 + 5^2 = 61
//  6^2 + 1^2 = 37  <--- pay attention
//  3^2 + 7^2 = 58
//  5^2 + 8^2 = 89
//  8^2 + 9^2 = 145
//  1^2 + 4^2 + 5^2 = 42
//  4^2 + 2^2 = 20
//  2^2 + 0^2 = 4
//  4^2 = 16
//  1^2 + 6^2 = 37  ---> here we have a loop

// Time complexity : O(log(n))
// Space complexity : O(log(n))
func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	// Space complexity goes mostly from here
	sums := make(map[int]struct{})

	for {
		// once next number goes below 999 it's impossible for it to go back to number with more than 3 digits
		if n = isHappyGetNext(n); n == 1 { // O(log(n))
			return true
		}

		// in the loop
		if _, ok := sums[n]; ok {
			return false
		}

		sums[n] = struct{}{}
	}
}

// Time complexity : O(log(n))
// Space complexity : O(1)
func isHappyFloyd(n int) bool {
	tortoise := n
	hare := isHappyGetNext(n)

	// Floyd's Cycle
	for hare != 1 && tortoise != hare {
		tortoise = isHappyGetNext(tortoise)
		hare = isHappyGetNext(isHappyGetNext(hare))
	}

	return hare == 1
}

// the largest next number we could get for each number of digits is:
//
// Digits Largest       Next
// 1      9             81
// 2      99            162
// 3      999           243
// 4      9999          324
// 10     2147483647    260   <-- max int32
func isHappyGetNext(n int) int {
	// sum of the squares of its digits
	var sum int

	// There are roughly log10(n) digits in x, We reduce a number of iterations by 10 - O(log(n))
	for n > 0 {
		// getting the last digit (example 19 % 10 = 9)
		lastDigit := n % 10
		// leftover of the original number (example 19 / 10 = 1)
		n = n / 10
		// adding square of the last digit
		sum += lastDigit * lastDigit
	}

	return sum
}
