package cmd

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/xbt573/goflagexample/cmd/handlers"
)

var (
	//go:embed usage
	Usage embed.FS

	fooCmd = flag.NewFlagSet("foo", flag.ExitOnError)
	barCmd = flag.NewFlagSet("bar", flag.ExitOnError)

	help = flag.Bool("help", false, "")
	h    = flag.Bool("h", false, "")

	fooMessage = fooCmd.String("message", "bar", "")
	barMessage = barCmd.String("message", "foo", "")
)

func init() {
	flag.Usage = func() {
		data, _ := Usage.ReadFile("usage/usage.txt")
		fmt.Print(string(data))
	}

	fooCmd.Usage = func() {
		data, _ := Usage.ReadFile("usage/foo.txt")
		fmt.Print(string(data))
	}

	barCmd.Usage = func() {
		data, _ := Usage.ReadFile("usage/bar.txt")
		fmt.Print(string(data))
	}
}

func Main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		handlers.FooHandler(handlers.FooOptions{
			Message: *fooMessage,
			Flags: *fooCmd,
		})

	case "bar":
		barCmd.Parse(os.Args[2:])
		handlers.BarHandler(handlers.BarOptions{
			Message: *barMessage,
			Flags: *barCmd,
		})

	default:
		flag.Usage()

		if !(*help || *h) {
			os.Exit(1)
		}
	}
}
