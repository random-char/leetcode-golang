package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	cases := map[int]string{
		4: "1211",
		5: "111221",
	}

	for n, expected := range cases {
		if got := solutions.CountAndSay(n); got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	}
}
