package main

import "fmt"

type Location struct {
	Latitude, Longitude float64
}

func main() {
	// Allocate an empty map with make
	m := make(map[string]Location)

	// Store the location of Heliox
	m["Heliox"] = Location{
		51.485945, 5.392263,
	}

	// Test and get the value of key "Heliox"
	loc, ok := m["Heliox"]
	fmt.Println(loc, ok)
}
