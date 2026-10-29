package main

import (
	"fmt"
	"os"

	"push-swap/parsing"
)

func main() {
	nombres, err := parsing.ParseArgs(os.Args[1:])
	fmt.Println(nombres, err)
}
