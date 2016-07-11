package main

import (
	"fmt"
)

// ByteCounter type counts the bytes written to it before discarding them
type ByteCounter int

// Write method of the ByteCounter type takes a byte slice of p bytes, and written
// the number of bytes written and any error.
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello %s", name)
	fmt.Println(c)
}
