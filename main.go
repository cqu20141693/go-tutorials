package main

import (
	"fmt"
	"github.com/cqu20141693/go-tutorials/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
