package main

import(
	"fmt"
)

var printf = fmt.Println

func main(){
	// workinf with functions
	name := "The IT howler 🐺"
	func_name(name)

	newName := [...] string{"solo", "mike", "carol", "jon"}
	age := [...] int{11,20,34,2}

	// name and age function
	for index :=0; index < len(newName); index++ {
		name_n_age(newName[index], age[index])
	}
}

// I like my functions to be below the main function
func func_name(name string){
	printf("The name is : ", name)
}

func name_n_age(name string, age int){
	printf(name, age)
}