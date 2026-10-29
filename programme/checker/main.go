package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ligne := scanner.Text()

		err = p.Apply(ligne)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
	if p.EstTrie() == true {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
}
