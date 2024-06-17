package solutions

import (
	"fmt"
	"strings"
)

var map1, map2 map[rune]int

type letterCounter []int

func NewLetterCounter(str string) *letterCounter {
	lc := make(letterCounter, 26)

	for _, r := range str {
		lc[r-'a'] += 1
	}

	return &lc
}

func (l *letterCounter) String() string {
	sb := strings.Builder{}
	for i, v := range *l {
		sb.WriteString(fmt.Sprintf("%d:%d,", 'a'+i, v))
	}

	return sb.String()
}

func groupAnagrams(strs []string) [][]string {
	groupsMap := make(map[string][]string)

	var lc *letterCounter
	var lcString string
	for _, str := range strs {
		lc = NewLetterCounter(str)
		lcString = lc.String()
		if _, ok := groupsMap[lcString]; ok {
			groupsMap[lcString] = append(groupsMap[lcString], str)
		} else {
			groupsMap[lcString] = []string{str}
		}
	}

	groups := make([][]string, len(groupsMap))
	i := 0
	for _, group := range groupsMap {
		groups[i] = group
		i += 1
	}
	return groups
}
