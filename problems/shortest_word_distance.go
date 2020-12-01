package problems

import "math"

// Given a list of words and two words word1 and word2, return the shortest
// distance between these two words in the list.
//
// Example:
// Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
//
// Input: word1 = “coding”, word2 = “practice”
// Output: 3
//
// Input: word1 = "makes", word2 = "coding"
// Output: 1
//
// Note:
// You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.

func shortestDistance(words []string, word1 string, word2 string) int {
	pos1, pos2 := -1, -1
	shortest := math.MaxInt32

	for i, word := range words {
		if word1 == word {
			pos1 = i
			if pos2 != -1 { // seen word2
				shortest = int(math.Min(float64(shortest), float64(pos1-pos2)))
			}
		} else if word2 == word {
			pos2 = i
			if pos1 != -1 { // seen word1
				shortest = int(math.Min(float64(shortest), float64(pos2-pos1)))
			}
		}

	}

	return shortest
}
