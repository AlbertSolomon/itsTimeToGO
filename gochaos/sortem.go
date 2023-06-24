package main

import(
	"fmt"
)

var print = fmt.Println

func main() {
	print("lets goooooo!")

	// simple bubble sort in go 

	collection := [...] int16 {1, 3, 10, 89, -7}
	collectionLength := len(collection)

	print(collectionLength)

	for itemx := 0; itemx < collectionLength - 1; itemx++ {
		for itemz := 0; itemz <collectionLength - 1; itemz++ {
			if collection[itemz] > collection[itemz + 1] {
				collection[itemz], collection[itemz+1] = collection[itemz + 1], collection[itemz]
			}
		}
	}

	print(collection)

}