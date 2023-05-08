// https://blog.jbowen.dev/2019/11/bash-command-line-completion-with-go/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

// test -x -y ok
// test -xy fail
// test --str test ok
func parseWithFlag() {
	x := flag.Bool("x", false, "-x")
	y := flag.Bool("y", false, "-y")
	// strs := flag.String("strs", false, "-y") // no str array
	str := flag.String("str", "", "-str")
	flag.Parse()

	fmt.Printf("  x: %v\n", *x)
	fmt.Printf("  y: %v\n", *y)
	// fmt.Printf("  strs: %v\n", *strs)
	fmt.Printf("  str: %v\n", *str)
}

// test -x -y ok
// test -xy ok
// test -varx -vary ok
// test -s test ok
// test --str test ok
// test -S test,test ok
// test --strs test,test ok
func parseWithGoFlags() {
	type options struct {
		X    bool     `short:"x" long:"varx" description:"bool var y"`
		Y    bool     `short:"y" long:"vary" description:"bool var y"`
		Strs []string `short:"S" long:"strs" env-delim:","`
		Str  string   `short:"s" long:"str"`
	}
	var opts options
	p := flags.NewParser(&opts, flags.Default&^flags.HelpFlag)
	_, err := p.Parse()
	if err != nil {
		fmt.Printf("fail to parse args: %v", err)
		os.Exit(1)
	}
	fmt.Printf("  x: %v\n", opts.X)
	fmt.Printf("  y: %v\n", opts.Y)
	fmt.Printf("  strs: %v\n", opts.Strs)
	fmt.Printf("  str: %v\n", opts.Str)
}

func main() {
	// parseWithFlag()
	parseWithGoFlags()
}
