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

	displayStuff("The type of number is :", reflect.TypeOf(number))
}