package iteration

import (
	"strings"
)

func Repeat(character string, repeats int) string {
	var output string
	for i := 0; i < repeats; i++ {
		output = output + character
	}
	return output
}

func CountNumberOf(target, sentence string) int {
	return strings.Count(sentence, target)
}
