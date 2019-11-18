package main

import "fmt"

func main() {
	// A classical for loop which counts from 1 to 3
	for n := 1; n < 4; n++ {
		fmt.Println("n =", n)
	}

	// Go doesn't have a while statement, for is used
	sum := 1
	for sum < 10 {
		sum++
	}
	fmt.Println("sum =", sum)
}
