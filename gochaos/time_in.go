package main

import (
	"fmt"
	"time"
)

var print = fmt.Printf

func main() {
	now := time.Now()

	print("The time is %v", now)
}