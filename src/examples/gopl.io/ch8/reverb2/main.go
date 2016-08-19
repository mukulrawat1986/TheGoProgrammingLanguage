// An echo server that tries to simulate the reverberations of a real echo.
// The server composes independent shouts.
package main

import (
	"time"
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // eg. connection aborted
			continue
		}
		go handleConn(conn) // create a new goroutine for each connection
	}
}

func handleConn(c net.Conn){
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1 * time.Second)
	}
	// Note: ignoring errors from input.Err()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
