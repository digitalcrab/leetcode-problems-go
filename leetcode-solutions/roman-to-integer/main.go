package main

// There are six instances where subtraction is used:
// 1. I can be placed before V (5) and X (10) to make 4 and 9.
// 2. X can be placed before L (50) and C (100) to make 40 and 90.
// 3. C can be placed before D (500) and M (1000) to make 400 and 900.
func romanToInt(s string) (sum int) {
	// store the mapping between roman and decimal
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// keep track of the last processed char
	var lastCh byte

	// loop from the back
	for i := len(s) - 1; i >= 0; i-- { // Time: O(n)
		// current char
		ch := s[i]

		// get the value
		v := values[ch] // O(1)

		// apply minus rule
		if (ch == 'I' && (lastCh == 'V' || lastCh == 'X')) ||
			(ch == 'X' && (lastCh == 'L' || lastCh == 'C')) ||
			(ch == 'C' && (lastCh == 'D' || lastCh == 'M')) {
			sum -= v
		} else {
			sum += v
		}

		// update last calc
		lastCh = ch
	}

	return sum
}
