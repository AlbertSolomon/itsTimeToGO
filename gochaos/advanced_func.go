package main 

import (
	"fmt"
)

var printf = fmt.Println


func main() {
	//printf("the percentages of the function are: ", percento())
	low, high := values()
	printf("the lower percent is:", lpercento(low))
	printf("the higher percent is:", hpercento(high))
}

// from GO by example 
func values()(int , int) {
	return 30, 70
}

func lpercento(low int) int {
	low = (low * 100) / 100
	return low
}

func hpercento(high int) int {
	high = (high * 100) / 100
	return high
}