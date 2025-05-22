package main

import "fmt"

type ValidWordAbbr struct {
	data map[string]map[string]int
}

func Constructor(dictionary []string) ValidWordAbbr {
	data := make(map[string]map[string]int)
	for _, s := range dictionary {
		a := abbr(s)
		// create a map
		if data[a] == nil {
			data[a] = make(map[string]int)
		}
		// save the word and how many times
		data[a][s]++
	}
	return ValidWordAbbr{data: data}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	a := abbr(word)
	words, found := this.data[a]
	// There is no entry in dictionary whose abbreviation is equal to `word` abbreviation.
	if !found {
		return true
	}

	// For any entry in dictionary whose abbreviation is equal to `word` abbreviation, that entry and `word` are the same.
	if words[word] > 0 && len(words) < 2 {
		return true
	}

	return false
}

// creates abbreviation
func abbr(s string) string {
	if len(s) <= 2 {
		return s
	}
	return fmt.Sprintf("%s%d%s", s[:1], len(s)-2, s[len(s)-1:])
}

func main() {

}
