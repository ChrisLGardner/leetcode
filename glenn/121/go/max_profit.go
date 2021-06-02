package glenn

func maxProfit(prices []int) int {
	mxProfit := 0
	beginVal := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < beginVal {
			beginVal = prices[i]
		} else if prices[i] - beginVal > mxProfit {
			mxProfit = prices[i] - beginVal
		}
	}

	return mxProfit;
}

// Runtime: 120 ms, faster than 96.40% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 8.3 MB, less than 97.00% of Go online submissions for Best Time to Buy and Sell Stock.
