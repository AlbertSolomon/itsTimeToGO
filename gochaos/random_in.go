package main

import (
	"fmt"
	"math/rand"
	"time"
)

var print = fmt.Println

func main() {
	seedSeconds := time.Now().Unix()
	rand.Seed(seedSeconds)
	
	randomNamber := rand.Intn(2020) + 1
	print("Random number:",randomNamber)
}