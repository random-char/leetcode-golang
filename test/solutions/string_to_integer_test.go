package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	cases := map[string]int{
		"9223372036854775808": 2147483647,
	}

	for s, expected := range cases {
		if got := solutions.MyAtoi(s); got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	}
}
