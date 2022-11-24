package reverse

import "strings"

func ReverseWords(input string) string {
	reversedSentenceSlice := []string{}
	wordSlice := strings.Split(input, " ")
	for j := 0; j < len(wordSlice); j++ {
		reversedWord := ""
		for i := len(wordSlice[j]) - 1; i >= 0; i-- {
			reversedWord = reversedWord + string(wordSlice[j][i])
		}
		reversedSentenceSlice = append(reversedSentenceSlice, reversedWord)
	}
	return strings.Join(reversedSentenceSlice, " ")
}
