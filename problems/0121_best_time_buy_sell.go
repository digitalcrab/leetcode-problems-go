package problems

// BestTimeToBuyAndSellStock121 solves https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
//
// You are given an array prices where prices[i] is the price of
// a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one
// stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction.
// If you cannot achieve any profit, return 0.
//
// Constraints:
// 1 <= prices.length <= 10^5
// 0 <= prices[i] <= 10^4
//
// Space complexity: O(1) because we only use two variables
// Time complexity: O(n) where n is the length of prices
func BestTimeToBuyAndSellStock121(prices []int) int {
	// start with the first price as the minimum price
	minPrice := prices[0]

	// start with 0 as the maximum profit, because we can't sell on the first day
	// and event if we do it, the profit will be 0
	maxProfit := 0

	for _, price := range prices[1:] {
		// if the current price is less than the minimum price, then update the minimum price
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			// if the current price is greater than the minimum price, then calculate the profit
			// and update the maximum profit
			maxProfit = profit
		}
	}

	return maxProfit
}
