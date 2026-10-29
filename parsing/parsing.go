package parsing

import (
	"errors"
	"strconv"
	"strings"
)

// ParseArgs renvoie une erreur si un morceau n'est pas un entier ou s'il apparaît deux fois.
func ParseArgs(args []string) ([]int, error) {
	var nombres []int
	morceaux := strings.Fields(strings.Join(args, " "))
	dejaVu := make(map[int]bool)

	for _, morceau := range morceaux {
		n, err := strconv.Atoi(morceau)
		if err != nil {
			return nil, err
		}
		if dejaVu[n] {
			return nil, errors.New("doublon")
		}
		dejaVu[n] = true
		nombres = append(nombres, n)
	}
	return nombres, nil
}
