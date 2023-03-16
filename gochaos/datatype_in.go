package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var displayStuff = fmt.Println

func main() {
	//TODO working wuth Data typrs

	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""

	const PI float32 = 3.14

	number := false
	number1 := 1
	number2 := 2.9
	number3 := "sup"

	displayStuff("The type of number is :", reflect.TypeOf(number))
	displayStuff("The type of number1 is :", reflect.TypeOf(number1))
	displayStuff("The type of number2 is :", reflect.TypeOf(number2))
	displayStuff("The type of number3 is :", reflect.TypeOf(number3))
	displayStuff("The type of PI is :", reflect.TypeOf(PI))

	// converting to INT
	number4 := int(number2)
	if reflect.TypeOf(number4) == reflect.TypeOf(number1) {
		displayStuff("number4 was conveted to int", number4)
	}

	// Lets convert strings to Integers
	number5 := "1200000000000"
	number6, err := strconv.Atoi(number5) // trying to handle the error
	displayStuff("converting strings to integers", number6, err, reflect.TypeOf(number6))
}
