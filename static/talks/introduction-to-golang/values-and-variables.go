package main

import "fmt"

func main() {
	// Concatenate and calculate
	fmt.Println("foo" + "bar")
	fmt.Println("12.3 * 2 =", 12.3*2)

	// Variables always have a initialized value
	var a uint
	fmt.Println("a =", a)

	// Multiple variables with multiple types
	var b, c = 10.5, "chicken"
	fmt.Printf("b(%v): %T, c(%v): %T\n", b, b, c, c)

	// Single variable declare and assign
	d := "var d"
	fmt.Println(d)
}
