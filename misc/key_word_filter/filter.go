package key_word_filter

import "strings"

func FilterKeyWords(words string) string {
	key_words := keyWords()

	for _, val := range key_words {
		if strings.Contains(words, val) {
			words = strings.Replace(words, val, "", -1)
		}
	}

	return words
}
