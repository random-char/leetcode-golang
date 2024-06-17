package solutions

func firstMissingPositive(nums []int) int {
	missing := -(len(nums) + 1)
	for i, num := range nums {
		if num < 1 || num > len(nums) {
			nums[i] = missing
		} else {
			nums[i] = -num
		}
	}

	for _, num := range nums {
		missing = max(missing, num)
	}

	return -missing
}
