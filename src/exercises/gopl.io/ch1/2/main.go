// Modify the echo program to print the index and value of each of the
// arguments, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%d %s\n", i, os.Args[i])
	}
}
