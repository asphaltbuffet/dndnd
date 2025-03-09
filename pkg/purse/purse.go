package purse

import (
	"fmt"
	"time"

	"github.com/asphaltbuffet/dndnd/pkg/calc"
)

type Purse struct {
	Name  string `toml:"name"`
	Value int    `toml:"value"`
}

func NewPurse(name string) *Purse {
	// quick and dirty way to generate a unique ID
	if name == "" {
		name = fmt.Sprintf("%d", time.Now().Unix())
	}

	p := &Purse{Name: name}

	return p
}

func (p *Purse) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Value)
}

func (p *Purse) Add(coins ...Coinage) {
	for _, c := range coins {
		c(p)
	}
}

func (p *Purse) Breakdown() CurrencyBreakdown {
	var b CurrencyBreakdown
	cp := p.Value

	b.Platinum, cp = cp/calc.PlatinumValue, cp%calc.PlatinumValue
	b.Gold, cp = cp/calc.GoldValue, cp%calc.GoldValue
	b.Electrum, cp = cp/calc.ElectrumValue, cp%calc.ElectrumValue
	b.Silver, cp = cp/calc.SilverValue, cp%calc.SilverValue
	b.Copper = cp

	return b
}

func (p *Purse) Copper() int {
	return p.Breakdown().Copper
}

func (p *Purse) Silver() int {
	return p.Breakdown().Silver
}

func (p *Purse) Electrum() int {
	return p.Breakdown().Electrum
}

func (p *Purse) Gold() int {
	return p.Breakdown().Gold
}

func (p *Purse) Platinum() int {
	return p.Breakdown().Platinum
}
