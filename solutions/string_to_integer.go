package solutions

import (
	"math"
	"strings"
	"unicode"
)

func MyAtoi(s string) int {
	return myAtoi(s)
}

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if s == "" {
		return 0
	}

	isNegative := s[0] == '-'
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	multiplier := 1
	if isNegative {
		multiplier = -1
	}

	result := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			result = result*10 + multiplier*int(r-'0')

			if result > math.MaxInt32 {
				return math.MaxInt32
			} else if result < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}

	return result
}
