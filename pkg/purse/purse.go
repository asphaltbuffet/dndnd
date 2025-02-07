package purse

import (
	"fmt"

	"github.com/asphaltbuffet/dndnd/pkg/calc"
)

type Purse struct {
	value int

	Copper   int
	Silver   int
	Electrum int
	Gold     int
	Platinum int
}

func NewPurse(value int) *Purse {
	p := &Purse{value: value}

	p.Platinum, value = value/calc.Platinum, value%calc.Platinum
	p.Gold, value = value/calc.Gold, value%calc.Gold
	p.Electrum, value = value/calc.Electrum, value%calc.Electrum
	p.Silver, value = value/calc.Silver, value%calc.Silver
	p.Copper = value

	return p
}

func (p *Purse) String() string {
	return fmt.Sprintf(
		"CP: %5d <<>> SP: %5d <<>> EP: %5d <<>> GP: %5d <<>> PP: %5d\n",
		p.Copper, p.Silver, p.Electrum, p.Gold, p.Platinum,
	)
}

func (p *Purse) Split(n int) ([]*Purse, int) {
	remainder := p.value % n
	quotient := p.value / n

	purses := make([]*Purse, n)
	for i := range purses {
		purses[i] = NewPurse(quotient)
	}

	return purses, remainder
}
