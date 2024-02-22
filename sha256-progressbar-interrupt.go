package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

func main() {
	const file = "/dev/zero"
	//const file = "./a.txt"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		h := sha256.New()
		bar := progressbar.DefaultBytes(
			-1,
			"hash",
		)
		if _, err := io.Copy(io.MultiWriter(h, bar), f); err != nil {
			if !errors.Is(err, os.ErrClosed) {
				log.Fatal("io.Copy ", err)
			}
			// if stoped by closing
		}
		fmt.Printf("\n%x  %s\n", h.Sum(nil), file)
	}()
	//
	go func() {
		time.Sleep(time.Second * 5)
		f.Close() // make error for stoping io.Copy
	}()
	wg.Wait()
}
