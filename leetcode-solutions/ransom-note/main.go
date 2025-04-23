package main

import "fmt"

// Time Complexity: O(m) see explanation inside
// Space Complexity: O(m) where m is 26 english letters, we can say it's constant O{1)
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	// create a map of characters and how many times we have them in the magazine
	mapped := make(map[rune]int, len(magazine))
	for _, ch := range magazine { // O(m) where m is number of characters in the magazine
		mapped[ch]++
	}

	// time complexity seems to me O(n) where n is number of chars in the ransomNote
	// however we break ot of the loop as soon as we do not have this char in the magazine
	// so it's safe to say that complexity is O(m) bounded by the number of chars in the magazine
	for _, ch := range ransomNote {
		// if there is no chars left, as well as no value present
		if mapped[ch] == 0 {
			return false
		}

		// sub one used char
		mapped[ch]--
	}

	return true
}

func main() {
	fmt.Printf("example 1 = %t\n", canConstruct("a", "b"))
	fmt.Printf("example 2 = %t\n", canConstruct("aa", "ab"))
	fmt.Printf("example 3 = %t\n", canConstruct("aa", "aab"))
}
