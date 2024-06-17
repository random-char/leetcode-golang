package solutions

func ConvertToTitle(columnNumber int) string {
	return convertToTitle(columnNumber)
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func convertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}

	columnNumber--
	letter := letters[(columnNumber)%26]
	columnNumber /= 26

	return convertToTitle(columnNumber) + string(letter)
}
