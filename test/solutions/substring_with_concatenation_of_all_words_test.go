package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

func TestFindSubscting(t *testing.T) {
	// solutions.FindSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"})
	solutions.FindSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})
}
