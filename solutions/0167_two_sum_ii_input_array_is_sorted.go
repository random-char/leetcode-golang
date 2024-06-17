package solutions

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	sum := numbers[start] + numbers[end]
	for sum != target {
		if sum > target {
			end--
		} else {
			start++
		}
		sum = numbers[start] + numbers[end]
	}

	return []int{start + 1, end + 1}
}
