// https://blog.jbowen.dev/2019/11/bash-command-line-completion-with-go/
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	debug := flag.Bool("debug", false, "enable debugging messages")
	dur := flag.Duration("duration", 5*time.Second, "duration to sleep")
	// Perform command line completion if called from completion library
	complete()
	flag.Parse()
	if *debug {
		fmt.Printf("Running '%s' with parameters:\n", os.Args[0])
		fmt.Printf("  debug:    %v\n", *debug)
		fmt.Printf("  duration: %v\n", *dur)
	}

	time.Sleep(*dur)
}
