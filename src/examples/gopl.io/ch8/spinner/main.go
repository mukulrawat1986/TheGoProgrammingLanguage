// In the example, the main goroutine computes the 45th fibonacci number.
// Since its a terribly recursive algorithm, it runs for an appreciable amount
// of time, during which we would like to provide the user with a visual indication
// that the program is running.
package main

import (
	"time"
	"fmt"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}