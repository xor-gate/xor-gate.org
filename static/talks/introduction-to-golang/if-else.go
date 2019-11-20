package main

import (
	"fmt"
	"time"
)

func main() {
	// Variables declared inside an if are also available in the else
	if name, offset := time.Now().Zone(); offset == 3600 {
		fmt.Println("We are here at UTC + 1 hour! (name: ", name, ")")
	} else if offset == 7200 {
		fmt.Println("We are here at UTC + 2 hour! (name: ", name, ")")
	} else {
		fmt.Println("We are at", name)
	}
}
