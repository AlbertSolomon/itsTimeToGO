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
		displayStuff("You are %v which is less than 18, hence too young for this\n", age)
	}

	if age >= 10 && age <=16 {
		displayStuff("you must be in high school")
	}else {
		displayStuff("i just dont know")
	}
}