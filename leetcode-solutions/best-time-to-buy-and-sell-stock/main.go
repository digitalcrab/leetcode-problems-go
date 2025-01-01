package main

func maxProfit(prices []int) int {
	// just a basic check for the correct input
	if len(prices) < 2 {
		return 0
	}

	// here we're going to store the max profit
	var bestProfit int
	// my current best price if I buy on the first day
	bestPrice := prices[0]

	// we loop over all prices except the first day as we just bought the stock
	for _, p := range prices[1:] {
		// it would be nice we found cheaper price, so lets buy this day instead
		if p < bestPrice {
			bestPrice = p
		} else {
			// here is the price higher lets calc the profit
			profit := p - bestPrice
			// and update if we made a better profit ;)
			bestProfit = max(bestProfit, profit)
		}
	}

	return bestProfit
}
