package solutions

func Rotate(nums []int, k int) {
	rotate(nums, k)
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverseArray(nums)
	reverseArray(nums[:k])
	reverseArray(nums[k:])
}

func reverseArray(nums []int) {
	i, j := 0, len(nums)-1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
