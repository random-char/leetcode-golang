package solutions

func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)

	for i, candidate := range candidates {
		if candidate == target {
			results = append(results, []int{candidate})
		} else if candidate < target {
			recursiveResults := combinationSum(candidates[i:], target-candidate)
			if len(recursiveResults) != 0 {
				for _, recursiveResult := range recursiveResults {
					if total(recursiveResult)+candidate == target {
						results = append(results, append(recursiveResult, candidate))
					}
				}
			}
		}
	}

	return results
}

func total(ints []int) int {
	result := 0

	for _, i := range ints {
		result += i
	}

	return result
}
