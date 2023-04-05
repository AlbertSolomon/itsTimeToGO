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
	
	printf("executing a lazy function")
	lazyLooping()
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

func lazyLooping(){
	for i := 0; i < 5; i++ {
		if i == 3 {
			printf("OOOh, man, can't manage this,lazy me: ", i)
			break
		}
		printf("the lazy i is ", i)
	}
}