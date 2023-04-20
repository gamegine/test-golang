package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// WriteCounter counts the number of bytes written to it.
type WriteCounter struct {
	Total int64 // Total # of bytes transferred
}

// Write implements the io.Writer interface.
//
// Always completes and never returns an error.
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += int64(n)
	fmt.Printf("Read %d bytes for a total of %d\n", n, wc.Total)
	return n, nil
}

func main() {

	var src io.Reader    // Source file/url/etc
	var dst bytes.Buffer // Destination file/buffer/etc

	// Create some random input data.
	src = bytes.NewBufferString(strings.Repeat("Some random input data", 1000))

	// Wrap it with our custom io.Reader.
	src = io.TeeReader(src, &WriteCounter{})

	count, err := io.Copy(&dst, src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Transferred", count, "bytes")
}
