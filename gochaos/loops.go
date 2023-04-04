package main

import (
	"fmt"
)

var print = fmt.Println

func main(){
	for i := 0; i < 10; i++ {
		print("solo", i)
	}
}