package main

import (
	"fmt"
	"reflect"
)

var displayStuff = fmt.Println

func main() {
	//TODO working wuth Data typrs
	
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	number := false
	number1 := 1
	number2 := 2.1
	number3 := "sup"

	displayStuff("The type of number is :", reflect.TypeOf(number))
	displayStuff("The type of number1 is :", reflect.TypeOf(number1))
	displayStuff("The type of number2 is :", reflect.TypeOf(number2))
	displayStuff("The type of number3 is :", reflect.TypeOf(number3))
}