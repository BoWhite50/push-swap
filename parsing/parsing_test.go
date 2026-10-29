package parsing

import (
	"slices"
	"testing"
)

func TestParseArgs(t *testing.T) {
	cas := []struct {
		nom     string
		args    []string
		attendu []int
		erreur  bool
	}{
		{"un seul argument", []string{"3 2 1"}, []int{3, 2, 1}, false},
		{"plusieurs arguments", []string{"3", "2", "1"}, []int{3, 2, 1}, false},
		{"mélange des deux", []string{"3 2", "1"}, []int{3, 2, 1}, false},
		{"signes", []string{"-5 +4 0"}, []int{-5, 4, 0}, false},
		{"espaces en trop", []string{"  3   2 "}, []int{3, 2}, false},
		{"une seule valeur", []string{"42"}, []int{42}, false},
		{"argument vide", []string{""}, nil, false},
		{"pas un nombre", []string{"3 one 1"}, nil, true},
		{"signe tout seul", []string{"-"}, nil, true},
		{"nombre à virgule", []string{"1.5"}, nil, true},
		{"doublon", []string{"1 2 1"}, nil, true},
		{"doublon entre deux arguments", []string{"1", "1"}, nil, true},
		{"nombre trop grand", []string{"99999999999999999999"}, nil, true},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			nombres, err := ParseArgs(c.args)
			if (err != nil) != c.erreur {
				t.Fatalf("erreur = %v, attendu une erreur : %v", err, c.erreur)
			}
			if !slices.Equal(nombres, c.attendu) {
				t.Errorf("ParseArgs(%q) = %v, attendu %v", c.args, nombres, c.attendu)
			}
		})
	}
}
