package main

import "math"

// Time Complexity: O(n) where n is number of chars in the string
// Space Complexity: O(26*2) where we store 26 times 2 indexes, so we can say O(1)
func firstUniqChar(s string) int {
	// here we gonna store how many times we've seen the character
	// 26 is ok as we are doing only lowercase english characters
	seen := make([][]int, 26)

	for idx, ch := range s { // O(n)
		// get the index of a counter
		countIdx := ch - 'a'
		// store indexes, and we basically need to know only the first and the fact that it's unique, so no reason to store more
		if len(seen[countIdx]) < 2 {
			seen[countIdx] = append(seen[countIdx], idx)
		}
	}

	sm := math.MaxInt

	for _, times := range seen { // O(26) = O(1)
		// seen once?
		if len(times) > 1 || len(times) == 0 {
			continue
		}
		// get the smallest idx
		sm = min(times[0], sm)
	}

	if sm == math.MaxInt {
		return -1
	}

	return sm
}

func main() {

}
