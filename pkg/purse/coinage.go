package purse

import "github.com/asphaltbuffet/dndnd/pkg/calc"

type Coinage func(*Purse)

func AddCopper(n int) Coinage {
	return func(p *Purse) {
		p.Value += n * calc.CopperValue
	}
}

func SubtractCopper(n int) Coinage {
	return func(p *Purse) {
		p.Value -= n * calc.CopperValue
	}
}

func AddSilver(n int) Coinage {
	return func(p *Purse) {
		p.Value += n * calc.SilverValue
	}
}

func SubtractSilver(n int) Coinage {
	return func(p *Purse) {
		p.Value -= n * calc.SilverValue
	}
}

func AddElectrum(n int) Coinage {
	return func(p *Purse) {
		p.Value += n * calc.ElectrumValue
	}
}

func SubtractElectrum(n int) Coinage {
	return func(p *Purse) {
		p.Value -= n * calc.ElectrumValue
	}
}

func AddGold(n int) Coinage {
	return func(p *Purse) {
		p.Value += n * calc.GoldValue
	}
}

func SubtractGold(n int) Coinage {
	return func(p *Purse) {
		p.Value -= n * calc.GoldValue
	}
}

func AddPlatinum(n int) Coinage {
	return func(p *Purse) {
		p.Value += n * calc.PlatinumValue
	}
}

func SubtractPlatinum(n int) Coinage {
	return func(p *Purse) {
		p.Value -= n * calc.PlatinumValue
	}
}
