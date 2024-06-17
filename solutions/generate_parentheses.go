package solutions

func generateParenthesis(n int) []string {
	result := []string{}
	addParenthesis(&result, "", n, 0, 0)
	return result
}

func addParenthesis(result *[]string, parenthesisString string, n, open, closed int) {
	if closed == n {
		*result = append(*result, parenthesisString)
	}

	if open > closed {
		addParenthesis(result, parenthesisString+")", n, open, closed+1)
	}

	if open < n {
		addParenthesis(result, parenthesisString+"(", n, open+1, closed)
	}
}
