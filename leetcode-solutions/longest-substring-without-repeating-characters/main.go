package main

// Time Complexity: O(n) where n is number of chars in the string
// Space Complexity O(n)
func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	longest, start := 0, 0

	for end := 0; end < len(s); end++ {
		// check if we've seen the character already
		if prevIdx, found := seen[s[end]]; found {
			// if so we need to move start to it, as the whole substring up till prevIdx
			// makes no interest for us, we have to keep start as latest as possible
			start = max(start, prevIdx+1)
		}

		seen[s[end]] = end
		longest = max(longest, end-start+1)
	}

	return longest
}

func main() {

}
