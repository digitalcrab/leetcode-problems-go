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
	for i := 1; i < len(prices); i++ {
		p := prices[i]

		// it would be nice we found cheaper price, so lets buy this day instead
		if p < bestPrice {
			bestPrice = p
		} else {
			// if we reached the end of the list of next day price is going to drop low
			if i == len(prices)-1 || prices[i+1] < p {
				// sell and calc the profit
				bestProfit += p - bestPrice
				// buy again
				bestPrice = p
			}
		}
	}

	return bestProfit
}
