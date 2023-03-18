package main

import (
	"fmt"
	"unicode/utf8"
)

var printer = fmt.Printf

func main() {
	// lets work with runes and strings

	abcdef := "abcdef"
	characterRune := utf8.RuneCountInString(abcdef)

	print("abcdef character rune is %v:\n", characterRune)

}