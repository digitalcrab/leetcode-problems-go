package main

func longestCommonPrefix(strs []string) (prefix string) {
	// nothing given?
	if len(strs) == 0 {
		return ""
	}
	// one element is the prefix itself
	if len(strs) == 1 {
		return strs[0]
	}

	// the length of the prefix
	prefixLen := 1

main:
	for prefixLen <= len(strs[0]) { // O(n) where n is number of chars in a first word
		// read the next char
		currentPrefix := strs[0][:prefixLen]

		for _, s := range strs[1:] { // O(m) where m is number of strings in the input
			// there is not more strings as long as prefix
			if len(s) < len(currentPrefix) {
				break main
			}

			// not matching
			if s[:prefixLen] != currentPrefix {
				break main
			}
		}

		// move to the next char, and update prefix
		prefixLen++
		prefix = currentPrefix
	}

	return
}

// Time: O(n * m)
func longestCommonPrefixSets(strs []string) (prefix string) {
	// nothing given?
	if len(strs) == 0 {
		return ""
	}

	// for the first string build all possible prefixes
	common := generatePrefixes(strs[0], len(strs[0]))

	// loop over the rest of strings
	for _, s := range strs[1:] { // Time: O(n) where n is number of strings in `strs`
		// build all possible prefixes but no more that we already have
		prefixes := generatePrefixes(s, len(common)) // O(m) where m is number of common prefixes

		// keep intersection here
		intersection := make(map[string]bool)

		// common is probably going to be the smallest one
		for p := range common { // O(m) where m is number of common prefixes
			if prefixes[p] {
				intersection[p] = true
			}
		}

		// keep only intersection of them
		common = intersection
	}

	// now find the longest
	for p := range common { // O(m) where m is number of common prefixes
		if len(p) > len(prefix) {
			prefix = p
		}
	}

	return
}

// generatePrefixes creates all possible prefixes of the word, but no more than maxLen
// Time: O(n) where n is number of chars in the string
func generatePrefixes(s string, maxLen int) map[string]bool {
	prefixes := make(map[string]bool)
	for i := 1; i < min(len(s), maxLen)+1; i++ {
		prefixes[s[:i]] = true
	}
	return prefixes
}
