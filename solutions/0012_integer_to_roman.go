package solutions

import "strings"

var roman = []string{
	"M",
	"CM",
	"D",
	"CD",
	"C",
	"XC",
	"L",
	"XL",
	"X",
	"IX",
	"V",
	"IV",
	"I",
}
var ints = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}

	s := strings.Builder{}
	for p, i := range ints {
		if num >= i {
			count := num / i
			for c := 0; c < count; c++ {
				s.WriteString(roman[p])
			}

			num = num % i
		}
	}

	return s.String()
}
