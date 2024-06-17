package solutions

func Generate(numRows int) [][]int {
	return generate(numRows)
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for rowNumber := 0; rowNumber < numRows; rowNumber++ {
		if rowNumber == 0 {
			result[rowNumber] = []int{1}
			continue
		}

		prevRow := result[rowNumber-1]
		newRow := make([]int, rowNumber+1)

		for i, v := range prevRow {
			newRow[i] += v
			newRow[i+1] += v
		}

		result[rowNumber] = newRow
	}

	return result
}
