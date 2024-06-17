package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

type isMatchTestCase struct {
	s        string
	p        string
	expected bool
}

func TestIsMatch(t *testing.T) {
	testCases := []isMatchTestCase{
		{s: "aa", p: "a", expected: false},
		{s: "aa", p: "a*", expected: true},
		{s: "abcdefghi", p: "a.*defghi", expected: true},
	}

	for _, testCase := range testCases {
		if got := solutions.IsMatch(testCase.s, testCase.p); got != testCase.expected {
			t.Errorf("Expected %t, got %t", testCase.expected, got)
		}
	}
}
