package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Tuesday {
	// You can use commas in a case to cover multiple statements
	case time.Tuesday, time.Friday:
		// By default, go switch-cases don't fallthrough
		fallthrough
	case time.Saturday:
		fmt.Println("Today is the big day")
	default:
		fmt.Println("Its a regular day")
	}
}
