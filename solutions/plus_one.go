package solutions

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			digits[i]++
		}

		carry = digits[i] > 9
		digits[i] %= 10
	}

	if carry {
		digits = append([]int{1}, digits...)
	}

	return digits
}
