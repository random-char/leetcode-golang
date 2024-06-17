package solutions

func Combine(n int, k int) [][]int {
	return combine(n, k)
}

func combine(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	backtrack(n, k, 1, &[]int{}, &result)
	return result
}

func backtrack(n, k, start int, track *[]int, result *[][]int) {
	if len(*track) == k {
		toAppend := make([]int, k)
		for i, t := range *track {
			toAppend[i] = t
		}
		*result = append(*result, toAppend)
		return
	}

	for i := start; i <= n; i++ {
		*track = append(*track, i)
		backtrack(n, k, i+1, track, result)
		*track = (*track)[:len(*track)-1]
	}
}
