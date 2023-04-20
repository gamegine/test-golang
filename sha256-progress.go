package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// WriteCounter counts the number of bytes written to it.
type WriteCounter struct {
	read int64 // Total # of bytes transferred
	size int64 //
}

// Write implements the io.Writer interface.
//
// Always completes and never returns an error.
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.read += int64(n)
	fmt.Printf("Read %d bytes for a total of %d %.2f\n", n, wc.read, float64(wc.read)/float64(wc.size)*100)
	return n, nil
}

func main() {
	const file = "a.txt"

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		// Could not obtain stat, handle error
	}
	var wc WriteCounter
	wc.size = fi.Size()

	// TeeReader not support os.File
	var fread io.Reader
	fread = f
	// Wrap it with our custom io.Reader.
	fread = io.TeeReader(fread, &wc)

	h := sha256.New()
	if _, err := io.Copy(h, fread); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x  %s\n", h.Sum(nil), file)

	content, err := ioutil.ReadFile(file + ".sha256")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
