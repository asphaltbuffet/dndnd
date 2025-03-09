package calc

type CurrencyValue int

const (
	Invalid       CurrencyValue = iota
	CopperValue                 = 1
	SilverValue                 = 10
	ElectrumValue               = 50
	GoldValue                   = 100
	PlatinumValue               = 1000
)

type CurrencyFunc func(int) int

func Copper(cp int) int {
	return cp * CopperValue
}

func Silver(sp int) int {
	return sp * SilverValue
}

func Electrum(ep int) int {
	return ep * ElectrumValue
}

func Gold(gp int) int {
	return gp * GoldValue
}

func Platinum(pp int) int {
	return pp * PlatinumValue
}
