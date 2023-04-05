package main

import (
	"fmt"
)

var printf = fmt.Println

func main(){

	for i := 0; i < 4; i++ {
		printf("solo", i)
	}

	printf("executing a function")
	furtherLooping()
}

func furtherLooping(){

	for i := 0; i < 5; i++ {
		if i == 3 {
			print("skipping ", i)
			print("\n")
			continue
		}
		printf("i is ", i)
	}

}