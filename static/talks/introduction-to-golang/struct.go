package main

import "fmt"

// Create a struct to store person information
type Person struct {
	Name string
	Age  int
}

func main() {
	// Create and initialize a person
	p := Person{Name: "Jerry", Age: 30}

	// Print struct including property names
	fmt.Printf("%+v\n", p)
}
