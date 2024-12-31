package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	// no way to have `needle` in `haystack`
	if len(needle) > len(haystack) {
		return -1
	}
	if haystack == needle {
		return 0
	}

	// start exploring from the beginning of the haystack
	// the end of the exploration should end in a way that needle can still fit in the haystack
	for haystackIdx := 0; haystackIdx <= len(haystack)-len(needle); haystackIdx++ {
		needleIdx := 0

		// now lets move along the needle
		for needleIdx < len(needle) {
			// check if we have match of chars in the haystack and needle
			// if not then just stop exploration
			if haystack[haystackIdx+needleIdx] != needle[needleIdx] {
				break
			}
			needleIdx++
		}

		// if we were able to explore the whole needle then it's a first idx of it
		if needleIdx == len(needle) {
			return haystackIdx
		}
	}

	return -1
}
