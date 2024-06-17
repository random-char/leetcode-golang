package solutions

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

func longestPalindrome(s string) string {
	result := ""
	var j, foundL, foundR int
	var isPalindromeSoFar bool

	for i := 0; i < len(s); i++ {
		j = len(s) - 1
		l, r := i, j
		isPalindromeSoFar = false

		for l <= r {
			if s[l] != s[r] {
				j--
				l = i
				r = j
				isPalindromeSoFar = false
				continue
			}

			if !isPalindromeSoFar {
				isPalindromeSoFar = true
				foundL = l
				foundR = r
			}

			l++
			r--
		}

		if isPalindromeSoFar && (foundR-foundL)+1 > len(result) {
			result = s[foundL : foundR+1]
		}
	}

	return result
}
