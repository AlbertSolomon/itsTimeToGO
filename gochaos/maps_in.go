package main

import (
	"fmt"
)

var printf = fmt.Println

func main(){
	// maps 
	n := map[string] int{"potato": 30, "tomatoes":60}
	printf(n)

	makingAmap := make(map[string] int)
	makingAmap["potato"] = 50
	makingAmap["orange"] = 1

	printf("making a map ", makingAmap)

	// empty maps
}