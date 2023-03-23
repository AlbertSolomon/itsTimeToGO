package main

import (
	"fmt"
)

var  print = fmt.Println

func  main(){
	
	myslices := [] string {}
	myslices = [] string {"myslices", "sugarcoats"}
	// myslices[1] = "sugar" // this spits out an error

	array := [...] int {10,20, 303, 33, 30, 40, 70}
	newSlice := array[2:5]

	print("myslices", myslices)
	print("myslices from index: ", myslices[0])
	
	print("myslices's capacity: ", cap(myslices))
	print("myslices's length: ", len(myslices))

	print("the length of a new slice: ", len(newSlice))
	print("the capacity of a new slice is : ", cap(newSlice))
	print("a new slice: ", newSlice)
}