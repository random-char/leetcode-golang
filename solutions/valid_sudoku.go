package solutions

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

func IsValidSudoku(board [][]byte) bool {
	return isValidSudoku(board)
}

func isValidSudoku(board [][]byte) bool {
	return validRows(board) && validColumns(board) && validBloks(board)
}

func validRows(board [][]byte) bool {
	var current, total int
	var currentSymbol byte
	for row := 0; row < 9; row++ {
		total = 1
		for col := 0; col < 9; col++ {
			currentSymbol = board[row][col]
			if currentSymbol == '.' {
				continue
			}

			current = primes[int(currentSymbol)-49]
			if total > 0 && total%current == 0 {
				return false
			}

			total *= current
		}
	}

	return true
}

func validColumns(board [][]byte) bool {
	var current, total int
	var currentSymbol byte
	for col := 0; col < 9; col++ {
		total = 1
		for row := 0; row < 9; row++ {
			currentSymbol = board[row][col]
			if currentSymbol == '.' {
				continue
			}

			current = primes[int(currentSymbol)-49]
			if total > 0 && total%current == 0 {
				return false
			}

			total *= current
		}
	}

	return true
}

func validBloks(board [][]byte) bool {
	var current, total int
	var currentSymbol byte
	var row, col int

	for block := 0; block < 9; block++ {
		total = 1
		dx := block % 3
		dy := block / 3

		for cell := 0; cell < 9; cell++ {
			row = dy*3 + cell/3
			col = dx*3 + cell%3
			currentSymbol = board[row][col]
			if currentSymbol == '.' {
				continue
			}
			current = primes[int(currentSymbol)-49]

			if total > 0 && total%current == 0 {
				return false
			}

			total *= current
		}
	}

	return true
}
