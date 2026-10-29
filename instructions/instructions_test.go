package instructions

import (
	"slices"
	"testing"
)

// TestInstructions vérifie chaque instruction : on part de piles connues,
func TestInstructions(t *testing.T) {
	tests := []struct {
		nom     string
		depart  Piles
		action  func(p *Piles)
		attendu Piles
	}{
		// Exemples normaux
		{"sa", Piles{A: []int{2, 1, 3}}, (*Piles).Sa, Piles{A: []int{1, 2, 3}}},
		{"sb", Piles{B: []int{7, 8, 9}}, (*Piles).Sb, Piles{B: []int{8, 7, 9}}},
		{"ss", Piles{A: []int{2, 1}, B: []int{7, 8}}, (*Piles).Ss, Piles{A: []int{1, 2}, B: []int{8, 7}}},
		{"pa", Piles{A: []int{1, 3}, B: []int{2, 7}}, (*Piles).Pa, Piles{A: []int{2, 1, 3}, B: []int{7}}},
		{"pb", Piles{A: []int{2, 1, 3}, B: []int{7, 8}}, (*Piles).Pb, Piles{A: []int{1, 3}, B: []int{2, 7, 8}}},
		{"ra", Piles{A: []int{2, 1, 3, 6}}, (*Piles).Ra, Piles{A: []int{1, 3, 6, 2}}},
		{"rb", Piles{B: []int{7, 8, 9}}, (*Piles).Rb, Piles{B: []int{8, 9, 7}}},
		{"rr", Piles{A: []int{1, 2, 3}, B: []int{7, 8, 9}}, (*Piles).Rr, Piles{A: []int{2, 3, 1}, B: []int{8, 9, 7}}},
		{"rra", Piles{A: []int{2, 1, 3, 6}}, (*Piles).Rra, Piles{A: []int{6, 2, 1, 3}}},
		{"rrb", Piles{B: []int{7, 8, 9}}, (*Piles).Rrb, Piles{B: []int{9, 7, 8}}},
		{"rrr", Piles{A: []int{1, 2, 3}, B: []int{7, 8, 9}}, (*Piles).Rrr, Piles{A: []int{3, 1, 2}, B: []int{9, 7, 8}}},

		// Cas limites : piles vides ou trop petites, rien ne doit changer ni planter
		{"sa pile vide", Piles{}, (*Piles).Sa, Piles{}},
		{"sa un seul élément", Piles{A: []int{5}}, (*Piles).Sa, Piles{A: []int{5}}},
		{"pa avec b vide", Piles{A: []int{1, 2}}, (*Piles).Pa, Piles{A: []int{1, 2}}},
		{"pb avec a vide", Piles{B: []int{7}}, (*Piles).Pb, Piles{B: []int{7}}},
		{"ra un seul élément", Piles{A: []int{5}}, (*Piles).Ra, Piles{A: []int{5}}},
		{"rra pile vide", Piles{}, (*Piles).Rra, Piles{}},
	}

	for _, tt := range tests {
		t.Run(tt.nom, func(t *testing.T) {
			p := tt.depart
			tt.action(&p)
			if !slices.Equal(p.A, tt.attendu.A) || !slices.Equal(p.B, tt.attendu.B) {
				t.Errorf("obtenu A=%v B=%v, attendu A=%v B=%v", p.A, p.B, tt.attendu.A, tt.attendu.B)
			}
		})
	}
}

// TestExempleDuSujet fait l'exemple du sujet.
func TestExempleDuSujet(t *testing.T) {
	p := Piles{A: []int{2, 1, 3, 6, 8, 5}}
	p.Sa()
	p.Pb()
	p.Pb()
	p.Pb()
	p.Rb()
	p.Rrr()
	p.Pa()
	p.Pa()
	p.Pa()
	if !slices.Equal(p.A, []int{1, 2, 3, 5, 6, 8}) || len(p.B) != 0 {
		t.Errorf("obtenu A=%v B=%v, attendu A=[1 2 3 5 6 8] B=[]", p.A, p.B)
	}
}
