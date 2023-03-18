package main

import (
	"fmt"
	"strings"
)

var print = fmt.Printf

func main() {
	// Strings in Gooooooooo
	// some string manipulation

	socialMedia := "Tweater"
	replacer := strings.NewReplacer("a", "t")

	editedSocialMedia := replacer.Replace(socialMedia)
	print("the correct social media string is %v", editedSocialMedia)
}