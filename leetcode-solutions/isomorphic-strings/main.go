package main

// Time Complexity: O(n) where n is number of character in the string
// Space Complexity: O(n) where n is number of available ascii characters
func isIsomorphic(s string, t string) bool {
	// basic case if strings are different no mapping possible
	if len(s) != len(t) {
		return false
	}

	// here we gonna store the mapping between t and s
	memo := make(map[byte]byte)
	// here we store the fact that we used
	used := make(map[byte]bool)

	// loop over characters in a first string
	for i := range s {
		// get chars that we are working with
		sCh := s[i]
		tCh := t[i]

		// fetch where is the s mapped to t
		mappedT, found := memo[sCh]

		// no mapping yet, then save it and move on
		if !found && !used[tCh] {
			memo[sCh] = tCh
			used[tCh] = true
		} else if mappedT != tCh { // if mapped but not equal to what we see already
			return false
		}
	}

	return true
}

func main() {

}
