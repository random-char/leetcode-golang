package solutions

func candy(ratings []int) int {
	result := make([]int, len(ratings))
	for i := 0; i < len(result); i++ {
		result[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			result[i] = result[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && result[i] <= result[i+1] {
			result[i] = result[i+1] + 1
		}
	}

	total := 0
	for _, r := range result {
		total += r
	}

	return total
}
