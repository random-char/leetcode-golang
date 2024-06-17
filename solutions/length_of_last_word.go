package solutions

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	i := strings.LastIndexByte(s, ' ')
	return len(s) - i - 1
}
