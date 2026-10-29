package tri

import "push-swap/instructions"

func Trier(nombres []int) []string {
	p := instructions.Piles{A: nombres}
	if p.EstTrie() == true {
		return nil
	} else {
		return nil
	}
}
