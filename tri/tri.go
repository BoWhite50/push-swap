package tri

import "push-swap/instructions"

// Trier renvoie la liste des instructions qui trient nombres.
func Trier(nombres []int) []string {
	var coups []string
	p := instructions.Piles{A: nombres}
	if p.EstTrie() {
		return nil
	}
	if len(p.A) == 2 {
		jouer(&p, &coups, "sa")
	}
	return coups
}

// jouer fait l'instruction sur les piles et la note dans coups.
func jouer(p *instructions.Piles, coups *[]string, nom string) {
	p.Apply(nom)
	*coups = append(*coups, nom)
}
