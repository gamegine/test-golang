package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
)

func main() {
	// const file = "/dev/zero"
	const file = "./a.txt"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	var size int64
	if fi.Size() == 0 {
		size = -1
	} else {
		size = fi.Size()
	}

	h := sha256.New()
	bar := progressbar.DefaultBytes(
		size,
		"hash",
	)
	if _, err := io.Copy(io.MultiWriter(h, bar), f); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	fmt.Printf("%x  %s\n", h.Sum(nil), file)

	content, err := ioutil.ReadFile(file + ".sha256")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
