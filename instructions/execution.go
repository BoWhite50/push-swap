package instructions

import (
	"bufio"
	"io"
	"strings"
)

// Executer lit les instructions dans entree (une par ligne) et les applique sur les piles
func (p *Piles) Executer(entree io.Reader) error {
	scanner := bufio.NewScanner(entree)
	for scanner.Scan() {
		ligne := strings.TrimSuffix(scanner.Text(), "\r") // PowerShell ajoute \r en fin de ligne
		if ligne == "" {
			continue // une ligne vide (un \n en trop) est ignorée
		}
		if err := p.Apply(ligne); err != nil {
			return err
		}
	}
	return scanner.Err()
}
