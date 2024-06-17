package solutions

func uniqueOccurrences(arr []int) bool {
	occurencesCounts := make(map[int]int)
	for _, i := range arr {
		if _, ok := occurencesCounts[i]; ok {
			occurencesCounts[i] += 1
		} else {
			occurencesCounts[i] = 0
		}
	}

	occurances := make(map[int]struct{})
	for _, count := range occurencesCounts {
		if _, ok := occurances[count]; ok {
			return false
		}

		occurances[count] = struct{}{}
	}

	return true
}
