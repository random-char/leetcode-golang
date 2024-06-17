package solutions

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	swap := 1

	prevNum := nums[0]
	prevCount := 1

	for peek := 1; peek < len(nums); peek++ {
		nums[swap] = nums[peek]

		if nums[peek] == prevNum {
			prevCount++
		} else {
			prevNum = nums[peek]
			prevCount = 1
		}

		if prevCount < 3 {
			swap++
		}
	}

	return swap
}
