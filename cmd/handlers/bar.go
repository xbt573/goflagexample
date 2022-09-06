package handlers

import (
	"flag"
	"fmt"
	"os"
)

type BarOptions struct {
	Message string
	Flags   flag.FlagSet
}

func BarHandler(options BarOptions) {
	if options.Flags.NArg() > 1 {
		fmt.Printf("too many arguments: %v\n", options.Flags.Args())
		options.Flags.Usage()

		os.Exit(1)
	}

	var out *os.File

	if options.Flags.Arg(0) != "" {
		var err error

		out, err = os.Create(options.Flags.Arg(0))
		if err != nil {
			fmt.Printf("failed to open file: %v", err)
			os.Exit(1)
		}
	} else {
		out = os.Stdout
	}

	fmt.Fprintf(out, "bar %v\n", options.Message)
}
