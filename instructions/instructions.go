package instructions

type piles struct {
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
func (p *piles) Pa() {
	push(&p.B, &p.A)
}
func (p *piles) Pb() {
	push(&p.A, &p.B)
}
func (p *piles) Sa() {
	swap(&p.A)
}
func (p *piles) Sb() {
	swap(&p.B)
}
func (p *piles) Ra() {
	rotate(&p.A)
}
func (p *piles) Rb() {
	rotate(&p.B)
}
func (p *piles) Rra() {
	reverseRotate(&p.A)
}
func (p *piles) Rrb() {
	reverseRotate(&p.B)
}
func (p *piles) Ss() {
	swap(&p.A)
	swap(&p.B)
}
func (p *piles) Rr() {
	rotate(&p.A)
	rotate(&p.B)
}
func (p *piles) Rrr() {
	reverseRotate(&p.A)
	reverseRotate(&p.B)
}
