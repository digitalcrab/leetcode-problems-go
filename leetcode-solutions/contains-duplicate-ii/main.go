package main

// Time Complexity: O(n) where n is number of elements in the array
// Space Complexity: O(k) where k is number of k, in a worst case it might be as big as n
func containsNearbyDuplicate(nums []int, k int) bool {
	// keep a record of seen items
	seen := make(map[int]bool)
	// look over the array
	for i, n := range nums {
		// if seen, then we have an answer
		if seen[n] {
			return true
		}
		// save that we have seen that element
		seen[n] = true

		// if the map grows more than k (size of our sliding windows)
		if len(seen) > k {
			// remove the first element that we've seen
			idx := i - k
			delete(seen, nums[idx])
		}
	}
	return false
}

func main() {

}
