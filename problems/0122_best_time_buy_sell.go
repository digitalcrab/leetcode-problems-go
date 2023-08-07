package problems

// BestTimeToBuyAndSellStock122 solves https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
//
// You are given an integer array prices where prices[i] is the price of a
// given stock on the i-th day.
//
// On each day, you may decide to buy and/or sell the stock.
// You can only hold at most one share of the stock at any time.
// However, you can buy it then immediately sell it on the same day.
//
// Find and return the maximum profit you can achieve.
//
// Constraints:
// 1 <= prices.length <= 3 * 10^4
// 0 <= prices[i] <= 10^4
//
// Space complexity: O(1) because we only use a few variables
// Time complexity: O(n * log n) where n is the length of prices (because of the nested loops)
func BestTimeToBuyAndSellStock122(prices []int) int {
	var profit int

	// simple solution with one loop
	for i := 0; i < len(prices)-1; i++ {
		if p := prices[i+1] - prices[i]; p > 0 {
			profit += p
		}
	}

	return profit
}

// BestTimeToBuyAndSellStock122LocalMinMax same as BestTimeToBuyAndSellStock122 but with local min/max
// Space complexity: O(1) because we only use a few variables
// Time complexity: O(n * log n) where n is the length of prices (because of the nested loops)
func BestTimeToBuyAndSellStock122LocalMinMax(prices []int) int {
	var totalProfit int

	for i := 0; i < len(prices); { // O(n)
		// search for local minimum (the next is smaller than one before)
		// we need to find the minimum to buy
		for i < len(prices)-1 && prices[i+1] <= prices[i] { // O(log n)
			i++
		}

		// if we are at the end of the array, then we can't buy
		if i == len(prices)-1 {
			break
		}

		buy := i
		i++ // increment i to start searching for the sell

		// search for local maximum (the next is greater than one before)
		for i < len(prices) && prices[i] >= prices[i-1] { // O(log n)
			i++
		}

		sell := i - 1

		totalProfit += prices[sell] - prices[buy]
	}

	return totalProfit
}
