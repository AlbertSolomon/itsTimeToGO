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
	number := 50
	network_available := false

	if age < 18 {
		displayStuff("You are %v which is less than 18, hence too young for this\n", age)
	}

	if age >= 10 && age <=16 {
		displayStuff("you must be in high schoo\n")
	} else if age == 10 {
		 displayStuff(" go watch ben 10")
	}else {
		displayStuff("i just dont know")
	}

	if number > 10 {
		displayStuff(" number is greater than 10 ")
	} else if number > 20 {
		displayStuff(" %v is 20 upwards ", number)
	}else {
		displayStuff("NUmber is less than %v ", number)
	}

	if network_available {
		if age > 10 {
			displayStuff(" %v is more than 10 years, adulthood awaits ..🚮", age)
		}
	}
}