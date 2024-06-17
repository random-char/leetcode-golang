package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

func TestRotateArray(t *testing.T) {
	solutions.Rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
