package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	// arrays in go are kinda weird, the syntax that is

	array := [4]int{1, 3, 4, 5}
	var array1 = [4]int{10, 30, 40, 50}

	print("array", array)
	print("array1", array1)

	number := array[1]
	array1[2] = 100

	sum := number + array1[2]
	print(" the sum is :", sum)

	// initializing specific elements
	array2 :=[5] int{}
	array2 = [5] int{1:10, 2:40}

	print("array 2:", array2)

	// Arrays with inferred length
	newArray :=[...] int {1, 2, 3}
	secondArray := [...] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	print("newArray:", newArray)
	print("secondArray:", secondArray)

	// Length of an array
	print("length:", len(secondArray))

}
