package solutions

import "math"

func Jump(nums []int) int {
	return jump(nums)
}

var lookup map[int]int

func jump(nums []int) int {
	lookup = map[int]int{}
	return jumpRec(nums, 0)
}

func jumpRec(nums []int, index int) int {
	if len(nums) < 2 {
		return 0
	}

	if found, ok := lookup[index]; ok {
		return found
	}

	minJumps := math.MaxInt - 1
	for jumpLen := 1; jumpLen <= nums[0] && jumpLen < len(nums); jumpLen++ {
		minJumps = min(minJumps, jumpRec(nums[jumpLen:], index+jumpLen))
	}

	lookup[index] = minJumps + 1
	return lookup[index]
}
