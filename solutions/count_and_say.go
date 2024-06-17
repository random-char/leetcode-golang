package solutions

import (
	"fmt"
	"strings"
)

func CountAndSay(n int) string {
	return countAndSay(n)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	sb := strings.Builder{}

	prevString := countAndSay(n - 1)
	var digit rune
	var count int

	for _, d := range prevString {
		if d == digit {
			count++
			continue
		}

		if count != 0 {
			sb.WriteString(fmt.Sprintf("%d%d", count, int(digit-'0')))
		}

		digit = d
		count = 1
	}

	sb.WriteString(fmt.Sprintf("%d%d", count, int(digit-'0')))
	return sb.String()
}
