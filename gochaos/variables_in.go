package main

import (
	"fmt"
)

var displayStuff = fmt.Println

func main() {
	// comments in go

	var variableName string = "Solomon"
	var v1, v2 = 1.2, 5.6
	var sum = v1 + v2


	displayStuff("My name is :", variableName)
	displayStuff("The sum is :", sum)
}