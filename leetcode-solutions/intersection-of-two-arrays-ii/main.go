package main

// Time Complexity: O(n+m) where n and m are number of elements in each array
// Space Complexity: O(m) where m is number of elements in a bigger array
func intersect(nums1 []int, nums2 []int) []int {
	// find out the smaller array
	small, big := nums1, nums2
	if len(nums1) > len(nums2) {
		small, big = nums2, nums1
	}

	// convert the bigger one into a map of a number and how many times i've seen it
	seen := make(map[int]int) // O(n)
	for _, n := range big {
		seen[n] += 1
	}

	var res []int

	// go over the small
	for _, n := range small { // O(m)
		// check if there was a one on big?
		if seen[n] > 0 {
			res = append(res, n)
			seen[n] -= 1
		}
	}

	return res
}

func main() {

}
