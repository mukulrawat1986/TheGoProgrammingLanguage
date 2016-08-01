package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {

	// initialize word counter
	*w = 0

	// create a bytes Buffer from the byte slice we are writing
	b := bytes.NewBuffer(p)

	// create a new scanner from this buffer
	s := bufio.NewScanner(b)

	// set up the split function
	s.Split(bufio.ScanWords)

	for s.Scan() {
		*w++
	}

	return len(p), nil
}

func main() {

	// create a WordCounter
	var w WordCounter

	// test out the WordCounter
	w.Write([]byte("Hello World"))
	// print out the number of words
	fmt.Println(w)

	// test out the WordCounter
	w.Write([]byte("Hello World, I am dog"))
	// print out the number of words
	fmt.Println(w)

	// test out the WordCounter
	w.Write([]byte("hello"))
	// print out the number of words
	fmt.Println(w)

}
