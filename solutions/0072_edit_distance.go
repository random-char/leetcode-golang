package solutions

func minDistance(word1 string, word2 string) int {
	lenW1, lenW2 := len(word1), len(word2)

	matrix := make([][]int, lenW1+1)
	for m := 0; m <= lenW1; m++ {
		matrix[m] = make([]int, lenW2+1)
	}

	for i := 1; i <= lenW1; i++ {
		matrix[i][0] = i
	}

	for j := 1; j <= lenW2; j++ {
		matrix[0][j] = j
	}

	var substitutionCost int
	for i := 1; i <= lenW1; i++ {
		for j := 1; j <= lenW2; j++ {
			if word1[i-1] == word2[j-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+substitutionCost,
			)
		}
	}

	return matrix[lenW1][lenW2]
}
