package solutions

func setZeroes(matrix [][]int) {
	forChangeRows := make([]bool, len(matrix))
	forChangeCols := make([]bool, len(matrix[0]))

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if matrix[x][y] == 0 {
				forChangeCols[y] = true
				forChangeRows[x] = true
			}
		}
	}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if forChangeRows[x] || forChangeCols[y] {
				matrix[x][y] = 0
			}
		}
	}
}
