package solutions

import (
	"strings"
)

func Convert(s string, n int) string {
	return convert(s, n)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	rowIndex := 0
	movingDown := true

	for _, c := range s {
		rows[rowIndex].WriteRune(c)

		if rowIndex == 0 {
			movingDown = true
		} else if rowIndex == numRows-1 {
			movingDown = false
		}

		if movingDown {
			rowIndex++
		} else {
			rowIndex--
		}
	}

	resultingSb := strings.Builder{}
	for _, sb := range rows {
		resultingSb.WriteString(sb.String())
	}

	return resultingSb.String()
}
