package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)

var displayStuff = fmt.Println
var statement = "Lets GOOO!💨👽"

func main() {
	displayStuff(statement)
	displayStuff("Whats your Name:")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		displayStuff("Lets Go!", name)
	} else {
		log.Fatal(err)
	}
	
}