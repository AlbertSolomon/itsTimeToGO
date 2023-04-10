package main

import (
	"fmt"
)

var printf = fmt.Println

func main() {
	printf("select date to excute this code ")

	day := 5

	switch day {
		case 1:
			printf("the day is monday")
		case 2:
			printf("the day is tuesday")
		case 3:
			printf("the day is wday")
		case 4:
			printf("the day is thursday")
		case 5:
			printf("the day is friiiii")
	}

	// another switch statement
	switch day {
		case 5,6:
			printf("yeeeeeey weeekend")
		case 1,2,3,4:
			printf("yeeeeeey we work")
	}

}
