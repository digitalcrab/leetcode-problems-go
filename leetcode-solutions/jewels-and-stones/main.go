package main

// Time Complexity: O(m+n) where m is length of jewels and n is length of stones
func numJewelsInStones(jewels string, stones string) int {
	// convert jewels into hash table so it's faster to search for
	searchJewels := make(map[rune]bool)
	for _, j := range jewels {
		searchJewels[j] = true
	}

	var num int

	for _, s := range stones {
		if searchJewels[s] {
			num++
		}
	}

	return num
}

func main() {

}
