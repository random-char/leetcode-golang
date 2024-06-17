package solutions

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	//go binary search
	l, r := 0, m*n-1
	for l <= r {
		mid := (r + l) / 2
		x, y := mid/n, mid%n
		if matrix[x][y] == target {
			return true
		}

		if matrix[x][y] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
