package main

// build a key, as a combination of distances between chars
func createKey(s string) string {
	var key string
	for i := 1; i < len(s); i++ {
		diff := (26 + s[i] - 'a' - s[i-1] - 'a') % 26
		key += string(diff + 'a') // + 'a' makes it just human-readable bringing back to the a-z range
	}
	return key
}

// Time Complexity: O(N) even if we loop a few times
// Space Complexity: O(n) by storing all elements again
func groupStrings(strings []string) [][]string {
	groups := make(map[string][]string)
	var ans [][]string

	for _, str := range strings {
		key := createKey(str)
		groups[key] = append(groups[key], str)
	}

	for _, group := range groups {
		ans = append(ans, group)
	}

	return ans
}

func main() {

}
