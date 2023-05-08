package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

type TestComplete string

func (t *TestComplete) Complete(match string) []flags.Completion {
	options := []string{"tata", "toto", "titi"}
	ret := make([]flags.Completion, 0, len(options))
	for _, o := range options {
		if strings.HasPrefix(o, match) {
			ret = append(ret, flags.Completion{
				Item: o,
			})
		}
	}
	return ret
}

func main() {
	type options struct {
		X    bool     `short:"x" long:"varx" description:"bool var y"`
		Y    bool     `short:"y" long:"vary" description:"bool var y"`
		Strs []string `short:"S" long:"strs" env-delim:"," default:"test"`
		// Str  string   `short:"s" long:"str" default:"test"`
		Str TestComplete `short:"s" long:"str" default:"test"`
	}

	var opts options
	p := flags.NewParser(&opts, flags.Default)
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
