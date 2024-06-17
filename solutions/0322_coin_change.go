package solutions

var dp []int

func coinChange(coins []int, amount int) int {
	dp = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for a := 1; a <= amount; a++ {
		for _, coin := range coins {
			if a-coin >= 0 {
				dp[a] = min(dp[a], dp[a-coin]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
