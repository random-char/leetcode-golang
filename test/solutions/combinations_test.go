package solutions_test

import (
	"leetcode-golang/solutions"
	"testing"
)

type testCaseCombinations struct {
	n        int
	k        int
	expected [][]int
}

func TestCombine(t *testing.T) {
	testCases := []testCaseCombinations{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}

	for _, testCase := range testCases {
		got := solutions.Combine(testCase.n, testCase.k)
		if !combinationsAreEqual(got, testCase.expected) {
			t.Errorf("Got:\n%+v, expected:\n%+v\n", got, testCase.expected)
		}
	}
}

func combinationsAreEqual(got, expected [][]int) bool {
	if len(got) != len(expected) {
		return false
	}

	for i, gotCombination := range got {
		expectedCombination := expected[i]

		if len(gotCombination) != len(expectedCombination) {
			return false
		}

		for j, num := range gotCombination {
			if expectedCombination[j] != num {
				return false
			}
		}
	}

	return true
}
