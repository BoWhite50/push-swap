package instructions

import "errors"

// A[0] correspond au sommet
type Piles struct {
	A []int
	B []int
}

func push(from *[]int, to *[]int) {
	if len(*from) > 0 {
		*to = append([]int{(*from)[0]}, *to...)
		*from = (*from)[1:]
	}
}

func swap(list *[]int) {
	if len(*list) > 1 {
		(*list)[0], (*list)[1] = (*list)[1], (*list)[0]
	}
}

func rotate(list *[]int) {
	if len(*list) > 1 {
		*list = append((*list)[1:], (*list)[0])
	}
}

func reverseRotate(list *[]int) {
	if len(*list) > 1 {
		*list = append([]int{(*list)[len(*list)-1]}, (*list)[:len(*list)-1]...)
	}
}
func (p *Piles) pa() {
	push(&p.B, &p.A)
}
func (p *Piles) pb() {
	push(&p.A, &p.B)
}
func (p *Piles) sa() {
	swap(&p.A)
}
func (p *Piles) sb() {
	swap(&p.B)
}
func (p *Piles) ra() {
	rotate(&p.A)
}
func (p *Piles) rb() {
	rotate(&p.B)
}
func (p *Piles) rra() {
	reverseRotate(&p.A)
}
func (p *Piles) rrb() {
	reverseRotate(&p.B)
}
func (p *Piles) ss() {
	swap(&p.A)
	swap(&p.B)
}
func (p *Piles) rr() {
	rotate(&p.A)
	rotate(&p.B)
}
func (p *Piles) rrr() {
	reverseRotate(&p.A)
	reverseRotate(&p.B)
}

// Apply applique l'instruction nom sur les piles.
// Elle renvoie une erreur si l'instruction n'existe pas.
func (p *Piles) Apply(nom string) error {
	switch nom {
	case "pa":
		p.pa()
	case "pb":
		p.pb()
	case "sa":
		p.sa()
	case "sb":
		p.sb()
	case "ss":
		p.ss()
	case "ra":
		p.ra()
	case "rb":
		p.rb()
	case "rr":
		p.rr()
	case "rra":
		p.rra()
	case "rrb":
		p.rrb()
	case "rrr":
		p.rrr()
	default:
		return errors.New("instruction inconnue")
	}
	return nil
}

// EstTrie dit si A est croissante et B est vide
func (p *Piles) EstTrie() bool {
	if len(p.B) != 0 {
		return false
	}

	for i := 1; i < len(p.A); i++ {
		if p.A[i-1] > p.A[i] {
			return false
		}
	}

	return true
}
