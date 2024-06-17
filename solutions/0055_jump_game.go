package solutions

func CanJump(nums []int) bool {
	return canJump(nums)
}

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	needJump := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= needJump {
			needJump = 1
		} else {
			needJump++
		}
	}

	return nums[0] >= needJump
}
