package instructions

import "errors"

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
func (p *Piles) Pa() {
	push(&p.B, &p.A)
}
func (p *Piles) Pb() {
	push(&p.A, &p.B)
}
func (p *Piles) Sa() {
	swap(&p.A)
}
func (p *Piles) Sb() {
	swap(&p.B)
}
func (p *Piles) Ra() {
	rotate(&p.A)
}
func (p *Piles) Rb() {
	rotate(&p.B)
}
func (p *Piles) Rra() {
	reverseRotate(&p.A)
}
func (p *Piles) Rrb() {
	reverseRotate(&p.B)
}
func (p *Piles) Ss() {
	swap(&p.A)
	swap(&p.B)
}
func (p *Piles) Rr() {
	rotate(&p.A)
	rotate(&p.B)
}
func (p *Piles) Rrr() {
	reverseRotate(&p.A)
	reverseRotate(&p.B)
}

func (p *Piles) Apply(nom string) error {
	switch nom {
	case "pa":
		p.Pa()
	case "pb":
		p.Pb()
	case "sa":
		p.Sa()
	case "sb":
		p.Sb()
	case "ss":
		p.Ss()
	case "ra":
		p.Ra()
	case "rb":
		p.Rb()
	case "rr":
		p.Rr()
	case "rra":
		p.Rra()
	case "rrb":
		p.Rrb()
	case "rrr":
		p.Rrr()
	default:
		return errors.New("instruction inconnue")
	}
	return nil
}

// EstTrie dit si A est croissante et B est vide.
func (p *Piles) EstTrie() bool {
	if p.B != nil{
		return false
	}

	for i := 1; i < len(p.A); i++ {
		if p.A[i-1]>p.A[i] {
			return	false
		}
	}

	return true
}