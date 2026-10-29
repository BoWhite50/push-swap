package main

import (
	"fmt"
	"os"
	"push-swap/parsing"
	"push-swap/tri"
)

// Programme push-swap : affiche les instructions qui trient la pile a
func main() {
	if len(os.Args) < 2 {
		return
	}
	nombres, err := parsing.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	coups := tri.Trier(nombres)
	for _, coup := range coups {
		fmt.Println(coup)
	}
}
