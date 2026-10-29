package tri

import (
	"math/rand"
	"slices"
	"testing"

	"push-swap/instructions"
)

// verifier joue les coups de Trier sur une nouvelle pile, comme le ferait checker,
// et renvoie le nombre de coups.
func verifier(t *testing.T, nombres []int) int {
	t.Helper()
	coups := Trier(slices.Clone(nombres))
	p := instructions.Piles{A: slices.Clone(nombres)}
	for _, coup := range coups {
		if err := p.Apply(coup); err != nil {
			t.Fatalf("instruction inconnue %q", coup)
		}
	}
	if !p.EstTrie() {
		t.Fatalf("%v n'est pas trié avec %v", nombres, coups)
	}
	return len(coups)
}

// permutations appelle f avec chaque ordre possible de list.
func permutations(list []int, debut int, f func([]int)) {
	if debut == len(list) {
		f(list)
		return
	}
	for i := debut; i < len(list); i++ {
		list[debut], list[i] = list[i], list[debut]
		permutations(list, debut+1, f)
		list[debut], list[i] = list[i], list[debut]
	}
}

func TestDejaTrie(t *testing.T) {
	for _, nombres := range [][]int{{}, {5}, {1, 2, 3}, {-3, 0, 8, 9, 12, 40}} {
		if coups := Trier(nombres); len(coups) != 0 {
			t.Errorf("Trier(%v) = %v, attendu aucun coup", nombres, coups)
		}
	}
}

func TestPetitesListes(t *testing.T) {
	maxCoups := map[int]int{2: 1, 3: 2, 4: 5, 5: 8}
	for taille, limite := range maxCoups {
		list := make([]int, taille)
		for i := range list {
			list[i] = i * 10
		}
		permutations(list, 0, func(nombres []int) {
			if n := verifier(t, nombres); n > limite {
				t.Errorf("%v : %d coups, maximum %d", nombres, n, limite)
			}
		})
	}
}

func TestExemplesDuSujet(t *testing.T) {
	// le sujet montre 6 coups pour le premier exemple et 7 pour le second
	if n := verifier(t, []int{4, 67, 3, 87, 23}); n > 6 {
		t.Errorf("4 67 3 87 23 : %d coups, le sujet en fait 6", n)
	}
	if n := verifier(t, []int{2, 1, 3, 6, 5, 8}); n > 7 {
		t.Errorf("2 1 3 6 5 8 : %d coups, le sujet en fait 7", n)
	}
}

func TestGrandesListes(t *testing.T) {
	for _, taille := range []int{6, 10, 100, 500} {
		for essai := 0; essai < 5; essai++ {
			nombres := rand.Perm(taille)
			for i := range nombres {
				nombres[i] -= taille / 2 // pour avoir aussi des négatifs
			}
			verifier(t, nombres)
		}
	}
}
