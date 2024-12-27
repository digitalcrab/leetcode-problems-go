package main

import "fmt"

func lengthOfLastWord(s string) int {
	idx, length := len(s)-1, 0

	// skip empty space if any
	for ; s[idx] == ' '; idx-- {

	}

	// count non-empty space
	for ; idx >= 0 && s[idx] != ' '; idx-- {
		length++
	}

	return length
}

func main() {
	l := lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Printf("Length of Last Word: %d\n", l)
}
