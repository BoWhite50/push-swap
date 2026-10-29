package main

import (
	"fmt"
	"os"

	"push-swap/instructions"
	"push-swap/parsing"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	nombres, err := parsing.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	p := instructions.Piles{A: nombres}

	if err := p.Executer(os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	if p.EstTrie() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
