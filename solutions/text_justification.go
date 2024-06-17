package solutions

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	justified := []string{}

	currentLineWordsLen := 0
	wordsOnLine := []string{}
	for _, word := range words {
		currentWordLen := len(word)

		if currentLineWordsLen+currentWordLen+len(wordsOnLine) > maxWidth {
			if len(wordsOnLine) == 1 {
				justified = AddLineLeftJustified(justified, wordsOnLine, maxWidth)
			} else {
				justified = AddLineCenterJustified(justified, wordsOnLine, currentLineWordsLen, maxWidth)
			}

			currentLineWordsLen = 0
			wordsOnLine = wordsOnLine[0:0]
		}

		currentLineWordsLen += currentWordLen
		wordsOnLine = append(wordsOnLine, word)
	}

	if len(wordsOnLine) != 0 {
		justified = AddLineLeftJustified(justified, wordsOnLine, maxWidth)
	}

	return justified
}

func AddLineCenterJustified(justified []string, wordsOnLine []string, currentLineWordsLen, maxWidth int) []string {
	spacesToAdd := maxWidth - currentLineWordsLen
	spacesForAll := spacesToAdd / (len(wordsOnLine) - 1)
	extraSpaces := spacesToAdd % (len(wordsOnLine) - 1)

	justifiedLine := ""
	var sb strings.Builder
	for i, wordOnLine := range wordsOnLine {
		sb.WriteString(wordOnLine)
		if i == len(wordsOnLine)-1 {
			break
		}

		for s := 0; s < spacesForAll; s++ {
			sb.WriteString(" ")
		}

		if extraSpaces != 0 {
			sb.WriteString(" ")
			extraSpaces--
		}
	}

	return append(justified, justifiedLine)
}

func AddLineLeftJustified(justified []string, wordsOnLine []string, maxWidth int) []string {
	leftJustifiedString := strings.Join(wordsOnLine, " ")
	for len(leftJustifiedString) < maxWidth {
		leftJustifiedString += " "
	}

	return append(justified, leftJustifiedString)
}
