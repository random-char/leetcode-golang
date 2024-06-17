package solutions

func firstUniqChar(s string) int {
	chars := make(map[rune]int)
	for i, c := range s {
		if _, ok := chars[c]; ok {
			chars[c] = -1
		} else {
			chars[c] = i
		}
	}

	result := -1
	for _, pos := range chars {
		if pos != -1 {
			if result == -1 {
				result = pos
			} else {
				result = min(result, pos)
			}
		}
	}

	return result
}
