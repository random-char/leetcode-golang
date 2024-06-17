package solutions

func FindSubstring(s string, words []string) []int {
	return findSubstring(s, words)
}

func findSubstring(s string, words []string) []int {
	result := make([]int, 0)

	wordLen := len(words[0])
	windowLen := wordLen * len(words)

	i, j := 0, windowLen

	lookupFreq := make(map[string]int)
	for _, word := range words {
		lookupFreq[word]++
	}

	for j <= len(s) {
		substrFreq := make(map[string]int)

		for k := i; k < j; k += wordLen {
			str := s[k : k+wordLen]
			if _, ok := lookupFreq[str]; !ok {
				break
			}

			substrFreq[str]++
		}

		if freqsEqual(&lookupFreq, &substrFreq) {
			result = append(result, i)
		}

		i++
		j++
	}

	return result
}

func freqsEqual(lookup, curr *map[string]int) bool {
	for s, i := range *lookup {
		if (*curr)[s] != i {
			return false
		}
	}
	return true
}
