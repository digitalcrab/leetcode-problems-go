package main

import "fmt"

// Time Complexity: O(n) where n is number of a new elements expected in the resulting array
// Space Complexity: O(1) if we do not count the answer size, but if we do then it's O(n)
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		number := i + 1
		str := ""

		if number%3 == 0 {
			str += "Fizz"
		}
		if number%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = fmt.Sprintf("%d", number)
		}

		res[i] = str
	}
	return res
}

func main() {
	result := fizzBuzz(3)
	fmt.Printf("Fizz Buzz: %v\n", result)
}
