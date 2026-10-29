// Package parsing transforme les arguments du programme en liste d'entiers.
// Il est partagé par checker et push-swap.
package parsing

import (
	"strconv"
	"strings"
)

func ParseArgs(args []string) ([]int, error) {
	var nombres []int

	morceaux := strings.Fields(strings.Join(args, " "))

	for i := 0; i < len(morceaux); i++ {
		n, err := strconv.Atoi(morceaux[i])
		if err != nil {
			return nil, err
		}
		nombres = append(nombres, n)
	}
	return nombres, nil
}
