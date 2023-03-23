package main

import (
	"fmt"
)

var  print = fmt.Println

func  main(){
	
	myslices := [] string {}
	myslices = [] string {"myslices", "sugarcoats"}
	// myslices[1] = "sugar" // this spits out an error

	print("myslices", myslices)
	print("myslices from index: ", myslices[0])
	
	print("myslices's capacity: ", cap(myslices))
	print("myslices's length: ", len(myslices))
}