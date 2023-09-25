package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		count := i + 1
		if strings.Contains("aeiouAEIOU", words[i][0:1]) {
			words[i] = words[i] + "ma"
		} else {
			firstChar := words[i][0:1]
			words[i] = words[i][1:] + firstChar + "ma"
		}
		words[i] = words[i] + strings.Repeat("a", count)
	}
	goatLatin := strings.Join(words, " ")
	return goatLatin
}

func main() {
	fmt.Printf("%v\n", toGoatLatin("I speak Goat Latin"))
}
