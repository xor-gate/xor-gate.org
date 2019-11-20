package main

import (
	"io"
	"os"
	"strings"
)

// Read data from r and write to os.Stdout
//  the io.Reader is an interface with only a Read method
func WriteToStdout(r io.Reader) {
	// Copy to an io.Writer (dest) from an io.Reader (source)
	//  this can be anything! File, network socket etc.
	io.Copy(os.Stdout, r)
}

func main() {
	// Create an io.Reader from a string and print it to os.Stdout
	sr := strings.NewReader("Hello string reader!")
	WriteToStdout(sr)
}
