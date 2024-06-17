package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

type Case struct {
	in      string
	out     string
	numRows int
}

func TestZigzagConversions(t *testing.T) {
	cases := []Case{
		{
			in:      "PAYPALISHIRING",
			out:     "PAHNAPLSIIGYIR",
			numRows: 3,
		},
	}

	for _, c := range cases {
		if result := solutions.Convert(c.in, c.numRows); result != c.out {
			t.Errorf("Expected %s, got %s", c.out, result)
		}
	}
}
