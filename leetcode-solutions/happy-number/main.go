package main

import (
	"fmt"
)

func isHappy(n int) bool {
	repeat := make(map[int]bool)

	for {
		if n == 1 {
			return true
		}

		// found a loop
		if repeat[n] {
			return false
		}

		// store the number to prevent the loops
		repeat[n] = true

		// store a new number after all ops
		newN := 0
		// reduce N on every step by 10
		for n > 0 {
			// individual digit
			digit := n % 10
			// make a sum for power of 2 per digit
			newN += digit * digit
			// save the rest
			n /= 10
		}
		// start over again
		n = newN
	}
}

func main() {
	fmt.Printf("Number 19 is happy: %t\n", isHappy(19))
}
