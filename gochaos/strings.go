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
	print("The length of the string is %v", len(editedSocialMedia))
	
	if strings.Contains(socialMedia, "w") {
		print(" \nThe string contains 'w'")
	}else{
		print(" \n the statement is false......")
	}

	emojiString := strings.Replace(editedSocialMedia, "e", "👾", -1) // -1 means search through the whole string
	print("the string becomes %v", emojiString)

	containsEsc := "\n ..next line \n"
	removeEsc := strings.TrimSpace(containsEsc)
	print("\n removining spaces %v", removeEsc)
}