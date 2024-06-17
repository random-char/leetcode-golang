package solutions

var digitsMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	results := make([]string, 0)

	if digits == "" {
		return nil
	}

	digit := string(digits[0:1])
	if len(digits) == 1 {
		return digitsMap[digit]
	}

	for _, letter := range digitsMap[digit] {
		for _, result := range letterCombinations(digits[1:]) {
			results = append(results, letter+result)
		}
	}

	return results
}
