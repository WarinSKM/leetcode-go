package maximumwordsfoundinsentences

import "strings"

func mostWordsFound(sentences []string) int {
	maxValue := 0
	for _, sentance := range sentences {
		word := strings.Count(sentance, " ") + 1
		if word > maxValue {
			maxValue = word
		}

	}
	return maxValue
}
