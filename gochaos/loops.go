package main

import (
	"fmt"
)

var printf = fmt.Println

func main(){
	
	for i := 0; i < 10; i++ {
		printf("solo", i)
	}
}