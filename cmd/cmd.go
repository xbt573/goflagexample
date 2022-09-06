package cmd

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/xbt573/goflagexample/cmd/handlers"
)

var (
	//go:embed usage/usage.txt
	usage    string
	//go:embed usage/foo.txt
	fooUsage string
	//go:embed usage/bar.txt
	barUsage string

	Usage = struct {
		Usage 	 string
		FooUsage string
		BarUsage string
	}{
		Usage: usage,
		FooUsage: fooUsage,
		BarUsage: barUsage,
	}

	fooCmd = flag.NewFlagSet("foo", flag.ExitOnError)
	barCmd = flag.NewFlagSet("bar", flag.ExitOnError)

	help = flag.Bool("help", false, "")
	h    = flag.Bool("h", false, "")

	fooMessage = fooCmd.String("message", "bar", "")
	barMessage = barCmd.String("message", "foo", "")
)

func init() {
	flag.Usage = func() {
		fmt.Printf(Usage.Usage)
	}

	fooCmd.Usage = func() {
		fmt.Printf(Usage.FooUsage)
	}

	barCmd.Usage = func() {
		fmt.Printf(Usage.BarUsage)
	}
}

func Main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		handlers.FooHandler(handlers.FooOptions{
			Message: *fooMessage,
		})

	case "bar":
		barCmd.Parse(os.Args[2:])
		handlers.BarHandler(handlers.BarOptions{
			Message: *barMessage,
		})

	default:
		flag.Usage()

		if !(*help || *h) {
			os.Exit(1)
		}
	}
}
