package main

import (
	"fmt"
)

var print = fmt.Println

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


}