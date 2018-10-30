package problems

import "math"

// Say you have an array for which the i-th element is the price of a given stock on day i.
// If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock),
// design an algorithm to find the maximum profit.
//
// Note that you cannot sell a stock before you buy one.
//
// Time complexity : O(n). Only a single pass is needed.
// Space complexity : O(1). Only two variables are used.
//
func MaxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}
