package main

import (
	"fmt"
)

var displayStuff = fmt.Printf

func main() {
	// conditions are next

	// Conditional Operators : > < >= <= == !=
	// logical operators : && || !

	age := 15

	if age < 18 {
		displayStuff("You are %v which is less than 18, hence too young for this", age)
	}
}