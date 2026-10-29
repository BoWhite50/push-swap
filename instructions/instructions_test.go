package instructions

import (
	"slices"
	"testing"
)

// TestInstructions vérifie chaque instruction : on part de piles connues,
// on applique l'instruction, puis on compare avec les piles attendues.
func TestInstructions(t *testing.T) {
	tests := []struct {
		nom     string
		depart  Piles
		action  func(p *Piles)
		attendu Piles
	}{
		// Exemples normaux
		{"sa", Piles{A: []int{2, 1, 3}}, (*Piles).sa, Piles{A: []int{1, 2, 3}}},
		{"sb", Piles{B: []int{7, 8, 9}}, (*Piles).sb, Piles{B: []int{8, 7, 9}}},
		{"ss", Piles{A: []int{2, 1}, B: []int{7, 8}}, (*Piles).ss, Piles{A: []int{1, 2}, B: []int{8, 7}}},
		{"pa", Piles{A: []int{1, 3}, B: []int{2, 7}}, (*Piles).pa, Piles{A: []int{2, 1, 3}, B: []int{7}}},
		{"pb", Piles{A: []int{2, 1, 3}, B: []int{7, 8}}, (*Piles).pb, Piles{A: []int{1, 3}, B: []int{2, 7, 8}}},
		{"ra", Piles{A: []int{2, 1, 3, 6}}, (*Piles).ra, Piles{A: []int{1, 3, 6, 2}}},
		{"rb", Piles{B: []int{7, 8, 9}}, (*Piles).rb, Piles{B: []int{8, 9, 7}}},
		{"rr", Piles{A: []int{1, 2, 3}, B: []int{7, 8, 9}}, (*Piles).rr, Piles{A: []int{2, 3, 1}, B: []int{8, 9, 7}}},
		{"rra", Piles{A: []int{2, 1, 3, 6}}, (*Piles).rra, Piles{A: []int{6, 2, 1, 3}}},
		{"rrb", Piles{B: []int{7, 8, 9}}, (*Piles).rrb, Piles{B: []int{9, 7, 8}}},
		{"rrr", Piles{A: []int{1, 2, 3}, B: []int{7, 8, 9}}, (*Piles).rrr, Piles{A: []int{3, 1, 2}, B: []int{9, 7, 8}}},

		// Cas limites : piles vides ou trop petites, rien ne doit changer ni planter
		{"sa pile vide", Piles{}, (*Piles).sa, Piles{}},
		{"sa un seul élément", Piles{A: []int{5}}, (*Piles).sa, Piles{A: []int{5}}},
		{"pa avec b vide", Piles{A: []int{1, 2}}, (*Piles).pa, Piles{A: []int{1, 2}}},
		{"pb avec a vide", Piles{B: []int{7}}, (*Piles).pb, Piles{B: []int{7}}},
		{"ra un seul élément", Piles{A: []int{5}}, (*Piles).ra, Piles{A: []int{5}}},
		{"rra pile vide", Piles{}, (*Piles).rra, Piles{}},
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
	p.sa()
	p.pb()
	p.pb()
	p.pb()
	p.rb()
	p.rrr()
	p.pa()
	p.pa()
	p.pa()
	if !slices.Equal(p.A, []int{1, 2, 3, 5, 6, 8}) || len(p.B) != 0 {
		t.Errorf("obtenu A=%v B=%v, attendu A=[1 2 3 5 6 8] B=[]", p.A, p.B)
	}
}
