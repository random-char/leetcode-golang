package solutions

var pesDpPrefix, pesDpSuffix map[int]int

func productExceptSelf(nums []int) []int {
	pesDpPrefix, pesDpSuffix = make(map[int]int), make(map[int]int)

	pesDpPrefix[0] = 1
	pesDpSuffix[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		pesDpPrefix[i] = pesDpPrefix[i-1] * nums[i]
	}

	for i := len(nums) - 1; i > 0; i-- {
		pesDpSuffix[i] = pesDpSuffix[i+1] * nums[i]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = pesDpPrefix[i] * pesDpSuffix[i]
	}

	return result
}
