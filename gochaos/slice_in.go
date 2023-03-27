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

	// lets make a slice 😉 
	slice := make([] int ,5, 10)
	slice2 := make([] int ,5)
	print("a slice: ", slice)
	print("slice2 with omitted capacity: ", slice2)

	print("slice capacity:", cap(slice))
	print("slice length:", len(slice))

	print("slice2 capacity:", cap(slice2))
	print("slice2 length:", len(slice2))

	slice = append(slice, 1, 3, 5, 6)
	print(slice)
	slice2 = append(slice2, 1, 3, 5, 6)
	print("slice 2: ", slice2)

	// more on go slices

}