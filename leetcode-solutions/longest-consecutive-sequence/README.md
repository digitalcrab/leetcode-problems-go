# Longest Consecutive Sequence

https://leetcode.com/problems/longest-consecutive-sequence/description/

## Problem Description

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

## Example 1

```text
Input: nums = [100,4,200,1,3,2]
Output: 4
```

**Explanation:** The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

## Example 2

```text
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
```

## Constraints

```text
0 <= nums.length <= 10**5
-10**9 <= nums[i] <= 10**9
```

## Benchmark results

```text
cpu: Apple M2 Pro
Benchmark_longestConsecutive-10                  2714517               420.6 ns/op
Benchmark_longestConsecutiveHash-10              2243179               539.5 ns/op
Benchmark_longestConsecutiveForest-10             345327              3393 ns/op
```