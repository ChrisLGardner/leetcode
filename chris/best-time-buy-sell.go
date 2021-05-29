package chris

func maxProfit(prices []int) int {

	profit := 0

	seen := map[int]bool{}

	for i, v := range prices {
		if _, ok := seen[v]; ok {
			continue
		}
		for j := (len(prices) - 1); j >= (i + 1); j-- {
			if (prices[j] - v) > profit {
				profit = prices[j] - v
			}
		}
		seen[v] = true
	}
	return profit
}
