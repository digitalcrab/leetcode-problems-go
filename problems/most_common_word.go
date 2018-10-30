package problems

// Given a paragraph and a list of banned words, return the most frequent word that is not in the list
// of banned words.  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.
//
// Words in the list of banned words are given in lowercase, and free of punctuation.
// Words in the paragraph are not case sensitive.
// The answer is in lowercase.
//
// Time Complexity: O(P + B), where P is the size of paragraph and B is the size of banned.
// Space Complexity: O(P + B), to store the count and the banned set.
//
func MostCommonWord(paragraph string, banned []string) string {
	words := make(map[string]int)

	// make a map of banned words
	bannedMap := make(map[string]struct{}, len(banned))
	for _, w := range banned {
		bannedMap[w] = struct{}{}
	}

	mostCommonCount := 0
	mostCommonWord := ""

	checkTheWord := func(word string) {
		// banned?
		if _, ok := bannedMap[word]; ok {
			return
		}

		count := words[word] + 1

		if count > mostCommonCount {
			mostCommonCount = count
			mostCommonWord = word
		}

		// update the counter
		words[word] = count
	}

	currentWord := make([]rune, 0)

	for _, r := range paragraph {
		if !isCharacter(r) {
			// we have a word
			if len(currentWord) > 0 {
				word := string(currentWord)

				// reset the word
				currentWord = []rune{}

				checkTheWord(word)
			}

			continue
		}

		// make it lowercase
		if 'A' <= r && r <= 'Z' {
			r += 'a' - 'A'
		}

		currentWord = append(currentWord, r)
	}

	if len(currentWord) > 0 {
		checkTheWord(string(currentWord))
	}

	return mostCommonWord
}

func isCharacter(c rune) bool {
	// a - 97
	// z - 122
	// A - 65
	// Z - 90
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
