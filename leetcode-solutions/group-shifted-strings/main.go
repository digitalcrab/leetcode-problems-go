package main

// build a key, as a combination of distances between chars
func createKey(s string) string {
	var key string

	for i := 1; i < len(s); i++ {
		// Step-by-step explanation for "az" and "ba":
		// For "az":
		// i = 1: diff = (26 + ('z' - 'a') - ('a' - 'a')) % 26 = (26 + 25 - 0) % 26 = 25
		// key += string(25 + 'a') = key += 'z'

		// For "ba":
		// i = 1: diff = (26 + ('a' - 'a') - ('b' - 'a')) % 26 = (26 + 0 - 1) % 26 = 25
		// key += string(25 + 'a') = key += 'z'

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
