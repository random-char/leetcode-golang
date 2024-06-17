package solutions

import (
	"slices"
)

func hIndex(citations []int) int {
	slices.Sort(citations)

	numCitations := len(citations)
	for i, c := range citations {
		res := numCitations - i
		if c >= res {
			return res
		}
	}

	return 0
}
