package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	cases := map[int]string{
		28:  "AB",
		52:  "AZ",
		701: "ZY",
		704: "AAB",
	}

	for in, out := range cases {
		if result := solutions.ConvertToTitle(in); result != out {
			t.Errorf("Expecting %s, got %s", out, result)
		}
	}
}
