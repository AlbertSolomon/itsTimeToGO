package main

import (
	"fmt"
)

var  print = fmt.Println

func  main(){
	
	myslices := [] string {}
	myslices = [] string {"myslices"}
	myslices[1] = "sugar" // this spits out an error

	print("myslices", myslices)
	print("myslices from index: ", myslices[0])
}