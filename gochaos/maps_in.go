package main

import (
	"fmt"
)

var printf = fmt.Println

func main(){
	// maps 
	n := map[string] int{"potato": 30, "tomatoes":60}
	printf(n)
}