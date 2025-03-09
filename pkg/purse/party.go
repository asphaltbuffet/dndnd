package purse

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
)

type Party struct {
	Name      string `toml:"name"`
	Unclaimed int    `toml:"unclaimed"`

	Players map[string]Purse `toml:"purse"`
}

func Load(fn string) (*Party, error) {
	pp := &Party{}

	fmt.Println("Loading party from", fn)

	_, err := toml.DecodeFile(fn, pp)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (pp *Party) Size() int {
	return len(pp.Players)
}

func (pp *Party) Split(total int) int {
	remainder := total % len(pp.Players)
	quotient := total / len(pp.Players)

	for _, v := range pp.Players {
		v.Value += quotient
	}

	return remainder
}

func (pp *Party) String() string {
	var b bytes.Buffer
	err := toml.NewEncoder(&b).Encode(pp)
	if err != nil {
		panic(err)
	}

	return b.String()
}
