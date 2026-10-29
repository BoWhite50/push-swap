// Package parsing transforme les arguments du programme en liste d'entiers.
// Il est partagé par checker et push-swap.
package parsing

import (
	"errors"
	"strconv"
	"strings"
)

func ParseArgs(args []string) ([]int, error) {
	var nombres []int

	morceaux := strings.Fields(strings.Join(args, " "))
	doublon := make(map[int]bool)

	for i := 0; i < len(morceaux); i++ {
		n, err := strconv.Atoi(morceaux[i])
		if err != nil {
			return nil, err
		}

		nombres = append(nombres, n)

		if doublon[n] == true {
			return nil, errors.New("doublon")
		} else {
			doublon[n] = true
		}
	}
	return nombres, nil
}
