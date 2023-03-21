package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	// arrays in go are kinda weird, the syntax that is

	array := [4] int{1,3,4,5}

	print("array", array)
}