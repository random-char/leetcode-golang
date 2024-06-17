package solutions

func minSubArrayLen(target int, nums []int) int {
	length := 0

	start := 0
	sum := 0

	var currLen int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= target {
			currLen = i - start + 1
			if length == 0 || currLen < length {
				length = currLen
			}

			sum -= nums[start]
			start++
		}
	}

	return length
}
