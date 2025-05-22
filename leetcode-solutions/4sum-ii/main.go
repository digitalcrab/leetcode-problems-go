package main

// Time Complexity: O(n^2)
// Spca complexity: O(n^2)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	abSums := make(map[int]int)

	// calculate first summ of first 2 arrays
	for a := 0; a < len(nums1); a++ {
		for b := 0; b < len(nums2); b++ {
			abSums[nums1[a]+nums2[b]]++
		}
	}

	var res int

	// loop over the rest of two arrays and find the nevagive sum of them
	for c := 0; c < len(nums3); c++ {
		for d := 0; d < len(nums4); d++ {
			cdSum := nums3[c] + nums4[d]

			if abSums[-cdSum] > 0 {
				// save how many of them we have
				res += abSums[-cdSum]
			}
		}
	}

	return res
}

func main() {

}
