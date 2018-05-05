package acronym

import (
	"strings"
)

// Abbreviate returns the acronym of given phrase
func Abbreviate(phrase string) string {
	words := strings.FieldsFunc(phrase, wordSplitter)
	acronym := make([]rune, len(words))
	for i, word := range words {
		acronym[i] = []rune(word)[0]
	}
	return strings.ToUpper(string(acronym))
}

func wordSplitter(r rune) bool {
	return r == ' ' || r == '-'
}
