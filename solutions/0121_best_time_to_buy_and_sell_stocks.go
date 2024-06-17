package solutions

func maxProfit(prices []int) int {
	minI, maxI, maxProfit := 0, 0, 0

	for i, p := range prices {
		if p < prices[minI] {
			maxProfit = max(maxProfit, prices[maxI]-prices[minI])
			minI, maxI = i, i
		}

		if p > prices[maxI] {
			maxI = i
		}
	}

	return max(maxProfit, prices[maxI]-prices[minI])
}
