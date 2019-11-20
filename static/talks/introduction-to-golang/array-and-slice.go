package main

import "fmt"

func main() {
	// An array is a fixed-size object
	wordarr := [4]byte{0xde, 0xad, 0xbe, 0xef}

	// Convert the array to a slice and typecast to string
	wordsli := wordarr[:]

	// Print the length and capacity of the slice (and hex)
	fmt.Printf("len(%d), cap(%d), %x\n", len(wordsli), cap(wordsli), string(wordsli))
}
