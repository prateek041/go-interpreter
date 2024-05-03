package main

import (
	"fmt"
	"os"

	"github.com/prateek041/go-interpreter/repl"
)

func main() {
	fmt.Println("Hello, welcome to the REPL")
	repl.Start(os.Stdout, os.Stdin)
}
