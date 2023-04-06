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

	printf("executing function with nestsed loops")
	nestemLoops()
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

func nestemLoops(){
	description := [...] string {"Big", "Tasty"}
	fruits := [...] string {"🍏", "🍑"}
	
	for i:=0 ; i<len(description); i++ {
		for fruit := 0; fruit < len(fruits); fruit++ {
			printf(description[i], fruits[fruit])
		}
	}
}