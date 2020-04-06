package problems

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/two-sum-less-than-k/
//
// Given an array A of integers and integer K, return the maximum S such that
// there exists i < j with A[i] + A[j] = S and S < K.
// If no i, j exist satisfying this equation, return -1.
//
// Example 1:
// Input: A = [34,23,1,24,75,33,54,8], K = 60
// Output: 58
// Explanation:
//  We can use 34 and 24 to sum 58 which is less than 60.
//
// Example 2:
// Input: A = [10,20,30], K = 15
// Output: -1
// Explanation:
//  In this case it's not possible to get a pair sum less that 15.
//
// Note:
//  - 1 <= A.length <= 100
//  - 1 <= A[i] <= 1000
//  - 1 <= K <= 2000
//

// Time Complexity: O(n*log(n)) + O(n) = O(n*log(n))
// Space Complexity: O(log(n))
func twoSumLessThanK(A []int, K int) int {
	// sort first, assume it takes - O(n*log(n))
	sort.Ints(A)

	i := 0
	j := len(A) - 1
	maxSum := math.MinInt32

	// moving while i < j, operation takes O(n)
	for i < j {
		newSum := A[i] + A[j]

		// if new sum is greater than K we should move right pointer
		// because we have only bigger elements from the right size
		if newSum >= K {
			j--
		} else {
			// otherwise move left side and update max
			maxSum = int(math.Max(float64(maxSum), float64(newSum)))
			i++
		}
	}

	if math.MinInt32 == maxSum {
		return -1
	}

	return maxSum
}
