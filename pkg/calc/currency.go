package calc

type Currency int

const (
	Invalid  Currency = iota
	Copper            = 1
	Silver            = 10
	Electrum          = 50
	Gold              = 100
	Platinum          = 1000
)
