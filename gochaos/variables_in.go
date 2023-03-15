package main

import (
	"fmt"
)

var displayStuff = fmt.Println

func main() {
	// comments in go

	var variableName string = "Solomon"
	var v1, v2 = 1.2, 5.6

	var sum = (v1 + v2)
	v3 := 7.8
	// its interesting that if i do this 👇🏾 with a variable it will be ok
	v3 = 23.7 // if i do this v3 := 23.7 it will raise an error, so interesting

	difference := sum - v3

	displayStuff("My name is :", variableName)
	displayStuff("The sum is :", sum)
	displayStuff("The difference is :", difference)
}
