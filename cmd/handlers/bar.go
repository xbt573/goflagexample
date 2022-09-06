package handlers

import "fmt"

type BarOptions struct {
	Message string
}

func BarHandler(options BarOptions) {
	fmt.Printf("bar %v\n", options.Message)
}
