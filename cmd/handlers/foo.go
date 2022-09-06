package handlers

import "fmt"

type FooOptions struct {
	Message string
}

func FooHandler(options FooOptions) {
	fmt.Printf("foo %v\n", options.Message)
}
