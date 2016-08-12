// clock2 is a TCP server that periodically writes the time.
package main

import (
	"net"
	"log"
	"io"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() // this call is blocking
		if err != nil {
			log.Print(err) // eg connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // eg. client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}