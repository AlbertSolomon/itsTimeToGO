package main

import(
	"fmt"
)

var printf = fmt.Println

func main(){
	// workinf with functions
	name := "The IT howler 🐺"
	func_name(name)

}

// I like my functions to be below the main function
func func_name(name string){
	printf("The name is : ", name)
}