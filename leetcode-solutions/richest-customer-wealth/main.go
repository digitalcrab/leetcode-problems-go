package main

import "fmt"

// Time Complexity: O(c * b), where c - is number of customers and b is max number of banks per customer
// Space Complexity: O(1) we use just one new int to store the result
func maximumWealth(accounts [][]int) int {
	richest := 0 // we know that it can not be less than 0 in the account

	for _ /*customer id*/, banks := range accounts {
		wealth := 0
		for _, amount := range banks {
			wealth += amount
		}
		if wealth > richest {
			richest = wealth
		}
	}

	return richest
}

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	wealth := maximumWealth(accounts)
	fmt.Printf("The richest customer is the customer that has the maximum wealth - %d\n", wealth)
}
