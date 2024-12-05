package main

import (
	"fmt"
	"maps"
	"slices"
)

// groupAnagrams returns anagrams grouped together.
// Time Complexity: N is number of strings, M is the longest string, O(N * M log M) as the largest operation
// Space Complexity: O(N * M) as we still create the group that potentially consists of all strings individually
func groupAnagrams(strs []string) [][]string {
	grouping := make(map[string][]string)
	for _, s := range strs { // O(N * M log M), as we call `sortedString` for each string
		// as anagram is a words that contains the same letters, it would make sense to sort letters in the word
		// and use it as a grouping key
		key := strKey(s)
		// add the string into it's group
		grouping[key] = append(grouping[key], s)
	}

	// let's put all groups back together
	groups := make([][]string, 0, len(grouping))
	for g := range maps.Values(grouping) { // O(N)
		groups = append(groups, g)
	}

	return groups
}

// strKey gives the sorted string by character back
// Time Complexity: O(n log(n)), where n the length of the string, based on the std library
// TODO: this function could be improved if we do not sort the string, but maybe build the array of char = used times
func strKey(s string) string {
	sorted := []rune(s)
	slices.Sort(sorted)
	return string(sorted)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	fmt.Printf("Anagram groups: %v\n", groups)
}
