package tri

import (
	"fmt"
	"slices"

	"push-swap/instructions"
)

// Trier renvoie la liste des instructions qui trient nombres.
func Trier(nombres []int) []string {
	var coups []string
	p := instructions.Piles{A: nombres}
	if p.EstTrie() {
		return nil
	}
	switch len(p.A) {
	case 2:
		jouer(&p, &coups, "sa")
	case 3:
		trier3(&p, &coups)
	case 4, 5, 6:
		trierPetit(&p, &coups)
	default:
		trierGrand(&p, &coups)
	}
	return coups
}

// jouer fait l'instruction sur les piles et la note dans coups.
func jouer(p *instructions.Piles, coups *[]string, nom string) {
	_ = p.Apply(nom) // les noms viennent de notre code, ils sont toujours valides
	*coups = append(*coups, nom)
}

// trier3 trie 3 nombres en 2 coups maximum :
// on envoie le plus grand en bas, puis on remet les 2 du haut dans l'ordre.
func trier3(p *instructions.Piles, coups *[]string) {
	switch posMax(p.A) {
	case 0:
		jouer(p, coups, "ra")
	case 1:
		jouer(p, coups, "rra")
	}
	if p.A[0] > p.A[1] {
		jouer(p, coups, "sa")
	}
}

// tousLesCoups contient les 11 instructions du sujet.
var tousLesCoups = []string{"sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr"}

// trierPetit trouve la suite de coups la plus courte pour 4 à 6 nombres on verifie en brute force.
func trierPetit(p *instructions.Piles, coups *[]string) {
	type etape struct {
		piles instructions.Piles
		coups []string
	}
	file := []etape{{*p, nil}}
	dejaVu := map[string]bool{fmt.Sprint(p.A, p.B): true}

	for len(file) > 0 {
		courant := file[0]
		file = file[1:]
		for _, nom := range tousLesCoups {
			suivant := copier(courant.piles) // une copie, pour ne pas abîmer courant
			suivant.Apply(nom)
			cle := fmt.Sprint(suivant.A, suivant.B)
			if dejaVu[cle] {
				continue
			}
			dejaVu[cle] = true
			suite := append(slices.Clone(courant.coups), nom)
			if suivant.EstTrie() {
				for _, coup := range suite {
					jouer(p, coups, coup)
				}
				return
			}
			file = append(file, etape{suivant, suite})
		}
	}
}

// copier renvoie une copie des deux piles sans modifier l'original
func copier(p instructions.Piles) instructions.Piles {
	var copie instructions.Piles
	copie.A = append(copie.A, p.A...)
	copie.B = append(copie.B, p.B...)
	return copie
}

// trierGrand trie dans les cas ou il y a plus de 6 nombres par paquets :
//  1- on envoie A dans B par paquets de petits nombres
//  2- on ramène toujours le plus grand de B en haut de A.
func trierGrand(p *instructions.Piles, coups *[]string) {
	p.A = rangs(p.A)
	paquet := 15
	if len(p.A) > 100 {
		paquet = 30
	}
	compteurPourB := 0
	for len(p.A) > 0 {
		rang := p.A[0]
		if rang <= compteurPourB {
			jouer(p, coups, "pb")
			jouer(p, coups, "rb")
			compteurPourB++
		} else if rang <= compteurPourB+paquet {
			jouer(p, coups, "pb")
			compteurPourB++
		} else {
			jouer(p, coups, "ra")
		}
	}
	for len(p.B) > 0 {
		monter(p, coups, posMax(p.B), len(p.B), "rb", "rrb")
		jouer(p, coups, "pa")
	}
}

// monter fait tourner une pile par le chemin le plus court
func monter(p *instructions.Piles, coups *[]string, pos int, taille int, r string, rr string) {
	if pos <= taille/2 {
		for ; pos > 0; pos-- {
			jouer(p, coups, r)
		}
	} else {
		for ; pos < taille; pos++ {
			jouer(p, coups, rr)
		}
	}
}

// posMax donne la position du plus grand nombre de list
func posMax(list []int) int {
	pos := 0
	for i := 1; i < len(list); i++ {
		if list[i] > list[pos] {
			pos = i
		}
	}
	return pos
}

// rangs remplace chaque nombre par sa place dans la liste triée
func rangs(nombres []int) []int {
	copie := slices.Clone(nombres)
	slices.Sort(copie)
	resultat := make([]int, len(nombres))
	for i, n := range nombres {
		resultat[i], _ = slices.BinarySearch(copie, n)
	}
	return resultat
}
