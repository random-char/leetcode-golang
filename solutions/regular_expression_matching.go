package solutions

func IsMatch(s, p string) bool {
	return isMatch(s, p)
}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	currentMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) || (currentMatch && isMatch(s[1:], p))
	}

	return currentMatch && isMatch(s[1:], p[1:])
}
