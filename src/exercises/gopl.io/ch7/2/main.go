package main

import (
	"bytes"
	"fmt"
	"io"
)

// countingWriter is an interface of type io.Writer, that wraps over any existing
// io.Writer.
// It contains a io.Writer and a counter to count the number of bytes
// written to the writer so far.
type countingWriter struct {
	count *int64
	w io.Writer
}

func (c countingWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	if err == nil {
		*c.count += int64(n)
	}
	return n, err
}

// CountingWriter is a function that given an io.Writer, returns a new
// Writer that wraps the original, and pointer to an int64 variable that at any
// given moment contains the number of bytes written to the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64){
	var counter int64
	newWriter := countingWriter{
		count: &counter,
		w: w,
	}

	return newWriter, newWriter.count
}

func main() {

	// create a new io.Writer
	b := bytes.NewBuffer([]byte(""))

	// pass it to CountingWriter
	newB, count := CountingWriter(b)

	// write something to newB
	newB.Write([]byte("Hello world"))

	// print out the number of bytes written
	fmt.Println(*count)
}