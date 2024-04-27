package main

import (
	"fmt"
	"interpreter/repl"
	"os"
)

func main() {
	fmt.Printf("Entered REPL shell; type a command below:\n")
	repl.Start(os.Stdin, os.Stdout)
}
