package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		ligne := strings.TrimSuffix(scanner.Text(), "\r")
		if ligne == "" {
			continue
		}
		if err := p.Apply(ligne); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	if p.EstTrie() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
