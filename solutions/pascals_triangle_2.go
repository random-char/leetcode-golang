package solutions

func getRow(rowIndex int) []int {
	row := []int{1}
	for i := 0; i < rowIndex; i++ {
		row = append(row, 0)
		for j := len(row) - 1; j > 0; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}

func GetRow(rowIndex int) []int {
	return getRow(rowIndex)
}
