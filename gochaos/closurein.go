package main

import (
	"fmt"
)

var printf = fmt.Println

func main() {
	callingintSequenceMutipletimes()
}

// from go by examples
func intSequence() func() int {
	seq := 0

	return func() int {
		seq ++
		return seq
	}
}


func callingintSequenceMutipletimes() {
	numberOftimes := 10
	nextInt := intSequence()
	
	for seq := 0; seq < numberOftimes; seq++ {
		printf(nextInt())
	}
}