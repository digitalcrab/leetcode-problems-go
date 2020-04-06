package problems

import "sort"

// https://leetcode.com/problems/group-anagrams/
//
// Given an array of strings, group anagrams together.
//
// Example:
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
//   [
//     ["ate","eat","tea"],
//     ["nat","tan"],
//     ["bat"]
//   ]
//
// Note:
//  - All inputs will be in lowercase.
//  - The order of your output does not matter.
//

type BytesSlice []byte

func (p BytesSlice) Len() int           { return len(p) }
func (p BytesSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p BytesSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func groupAnagrams(input []string) [][]string {
	res := make(map[string][]string)

	for _, str := range input {
		sorted := BytesSlice(str)
		sort.Sort(sorted)
		res[string(sorted)] = append(res[string(sorted)], str)
	}

	groups := make([][]string, 0, len(res))
	for k := range res {
		groups = append(groups, res[k])
	}

	return groups
}
