package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	cases := map[string]string{
		"abb":                 "bb",
		"babad":               "bab",
		"cbbd":                "bb",
		"aacabdkacaa":         "aca",
		"bacabab":             "bacab",
		"xaabacxcabaaxcabaax": "xaabacxcabaax",
	}

	for in, out := range cases {
		if found := solutions.LongestPalindrome(in); found != out {
			t.Errorf("Got %s, expecting %s", found, out)
		}
	}
}
