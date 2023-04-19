package main 

import (
	"fmt"
)

var printf = fmt.Println

func main() {
	printf("There sum is ", sums(1, 1005, 60, 79, 90))

	nums := [] int {50, 55, 200, 45, 7}
	printf("sum of numbers in nums array is -> ", sums(nums...))
}

func sums(numbers... int) int {
	printf("The numbers are -> ", numbers)
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}