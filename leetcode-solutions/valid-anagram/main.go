package main

import "fmt"

// Time Complexity: O(n) + O(m) where n and m are number of characters in both strings, and we loop over all of them
// Space Complexity: O(26) = O(1), as we use only space required for the extra map with chars and counts
func isAnagram(s string, t string) bool {
	// if they are different in length then for sure not anagram
	if len(s) != len(t) {
		return false
	}

	// first try building the map of characters and how many times they used in the `s`
	charCounts := make(map[rune]int)
	for _, ch := range s {
		charCounts[ch]++ // charCounts[ch] gives you back 0 if value not present (default value of missing item)
	}

	// now `t` need's to use all the chars from `s`
	for _, ch := range t {
		// if value is not present this is for sure is not anagram
		if _, exists := charCounts[ch]; !exists {
			return false
		}

		// mark as used
		charCounts[ch]--

		// and delete from the map
		if charCounts[ch] <= 0 {
			delete(charCounts, ch)
		}
	}

	// when nothing left then probably anagram
	return len(charCounts) == 0
}

func main() {
	fmt.Printf("anagram and nagaram = %t", isAnagram("anagram", "nagaram"))
}
