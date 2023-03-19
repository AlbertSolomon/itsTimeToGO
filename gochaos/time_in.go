package main

import (
	"fmt"
	"time"
)

var print = fmt.Printf

func main() {
	now := time.Now()

	print("The time is %v", now)
	print("\nThe Year is %v , Month is %v and the day is %v", now.Year(), now.Month(), now.Day())
	print("\nThe hour is %v , Minute is %v and the second is %v", now.Hour(), now.Minute(), now.Second())
	
}