package main

import "math"

func mapping(l []string) map[string]int {
	mapped := make(map[string]int, len(l))
	for i := range l {
		mapped[l[i]] = i
	}
	return mapped
}

// Time Complexity: O(n+m) where n and m are number of strings in the respectful lists
// Space Complexity: O(n) where n is number if string in the list1 (use for mapped)
func findRestaurant(list1 []string, list2 []string) []string {
	// if one of them empty then no common strings
	if len(list1) == 0 || len(list2) == 0 {
		return []string{}
	}

	// convert both to maps
	mapped1 := mapping(list1) // O(n)

	var result []string
	lowestSum := math.MaxInt

	// loop over the one of them
	for idx2, s := range list2 { // O(m)
		// find the index of a second thing
		idx1, found := mapped1[s] // O(1)

		// if the same not found just skip it
		if !found {
			continue
		}

		// calculated new sum
		sum := idx1 + idx2

		// we found new winner
		if sum < lowestSum {
			// reset the results
			result = []string{s}
			lowestSum = sum
		} else if sum == lowestSum { // if we found the same sum
			result = append(result, s)
		}
	}

	return result
}

func main() {

}
