package solutions

var numDecDp map[int]int

func numDecodings(s string) int {
	numDecDp = map[int]int{
		len(s): 1,
	}

	numDecRec(0, s)

	return numDecDp[0]
}

func numDecRec(i int, s string) int {
	if _, ok := numDecDp[i]; ok {
		return numDecDp[i]
	}

	if s[i] == '0' {
		return 0
	}

	res := numDecRec(i+1, s)
	if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] < '7')) {
		res += numDecRec(i+2, s)
	}
	numDecDp[i] = res

	return res
}
