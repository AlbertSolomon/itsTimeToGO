package main

import (
	"fmt"
)

var printf = fmt.Println

type Names struct {
	name string 
	age int
} 

func main(){
	var named1 Names
	named1.name = "solo"
	named1.age =100

	var name Names
	name.name = "captain"
	name.age = 10

	printf("name is:", name.name)
	printf("name's age is:", name.age)
	
	printf("named1's age is:", named1.name)
	printf("named1's age is:", named1.age)

}