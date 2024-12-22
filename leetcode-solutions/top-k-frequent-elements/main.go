package main

import "fmt"

// topKFrequent returns k top frequent elements
// Time Complexity: O(n) as we loop over elements (multiple times)
// Space Complexity: O(N) for the counter map and buckets storage
func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int) // remelts to count, memory O(n)
	for _, n := range nums {     // time O(n)
		counter[n]++
	}

	// create a buckets, where key is how many times we see the element and a value
	// is the list of elements, memory in a work case O(n)
	buckets := make([][]int, len(nums)+1) // +1 is because we can as many elements as len(num) and index in this case could be 1 more

	for elem, count := range counter { // time O(n)
		// add elements to the count bucket
		buckets[count] = append(buckets[count], elem)
	}

	// here is going to be the result, k most frequent elements
	var res []int

	// going backwards from the most frequent
	for i := len(buckets) - 1; i >= 0; i-- { // time O(n)
		// collected all we need, or even more
		if len(res) >= k {
			break
		}

		// there is nothing in this bucket
		if len(buckets[i]) == 0 {
			continue
		}

		// add all elements of that count (this might be more than we need, but thats OK)
		res = append(res, buckets[i]...)
	}

	return res[:k] // return only as much as we need, as we probably can have more
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Printf("Top %d elements: %v\n", k, result)
}
