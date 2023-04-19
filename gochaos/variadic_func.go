package main 

import (
	"fmt"
)

var printf = fmt.Println

func main() {
	printf("There sum is ", sums(1,1005,60,79,90))
}

func sums(numbers... int) int {
	printf("The numbers are -> ", numbers)
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}