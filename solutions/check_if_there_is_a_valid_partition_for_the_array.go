package solutions

func validPartition(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	n := len(nums)
	dp := []bool{true, false, nums[0] == nums[1]}

	for i := 2; i < n; i++ {
		currentDp := false

		currentDp = nums[i] == nums[i-1] && dp[1] ||
			nums[i] == nums[i-1] && nums[i] == nums[i-2] && dp[0] ||
			nums[i]-1 == nums[i-1] && nums[i]-2 == nums[i-2] && dp[0]

		dp[0], dp[1], dp[2] = dp[1], dp[2], currentDp
	}

	return dp[2]
}

func ValidPartition(nums []int) bool {
	return validPartition(nums)
}
