package main

import (
	"fmt"
)

var displayStuff = fmt.Println

func main() {
	// comments in go

	var variableName string = "Solomon"
	displayStuff("My name is :", variableName)
}