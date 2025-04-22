package main

import "fmt"

// Time Complexity: O(log n) where n is how bit is the input number, and as we move we are halving the number
// Space Complexity: O(1) no extra space used
func numberOfSteps(num int) int {
	var steps int
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}

func main() {
	steps := numberOfSteps(14)
	fmt.Printf("Number of steps: %d\n", steps)
}
