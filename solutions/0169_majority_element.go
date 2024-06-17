package solutions

func majorityElement(nums []int) int {
	counts := map[int]int{}

	for _, num := range nums {
		counts[num]++
	}

	majNum, maxCount := 0, 0
	for num, count := range counts {
		if count > maxCount {
			maxCount = count
			majNum = num
		}
	}

	return majNum
}
