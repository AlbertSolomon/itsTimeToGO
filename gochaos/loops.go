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

	printf("executing function with nested loops")
	nestemLoops()

	printf("executing function with range in a loop")
	loopsofrange()
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

var description = [...] string {"Big", "Tasty"}
var fruits = [...] string {"🍏", "🍑"}

func nestemLoops(){	
	for i := 0; i < len(description); i++ {
		for fruit := 0; fruit < len(fruits); fruit++ {
			printf(description[i], fruits[fruit])
		}
	}
}

// a different implementation using range 
// go giving me c vibes 🤣
func loopsofrange(){
	for index, value := range fruits {
		printf(index, value)
	}
}