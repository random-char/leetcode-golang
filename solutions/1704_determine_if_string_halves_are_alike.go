package solutions

import "strings"

func halvesAreAlike(s string) bool {
	vowels := "aeiouAEIOU"
	v1, v2 := 0, 0

	halfLen := len(s) / 2
	for i := 0; i < halfLen; i++ {
		if strings.Contains(vowels, string(s[i])) {
			v1 += 1
		}

		if strings.Contains(vowels, string(s[halfLen+i])) {
			v2 += 2
		}
	}

	return v1 == v2
}
