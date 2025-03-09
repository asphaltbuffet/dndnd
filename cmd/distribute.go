package cmd

import (
	"github.com/asphaltbuffet/dndnd/pkg/calc"
	"github.com/asphaltbuffet/dndnd/pkg/purse"
	"github.com/spf13/cobra"
)

var distCmd *cobra.Command

var (
	cp int
	gp int
	sp int
	ep int
	pp int

	equal bool

	fn string
)

func GetDistributeCommand() *cobra.Command {
	if distCmd == nil {
		distCmd = &cobra.Command{
			Use:          "distribute",
			Short:        "distribute wealth to the party",
			Args:         cobra.NoArgs,
			Aliases:      []string{"d", "dist", "split"},
			SilenceUsage: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				party, err := purse.Load(fn)
				if err != nil {
					return err
				}

				err = Distribute(party)
				if err != nil {
					return err
				}

				// TODO: save party back to file

				cmd.Println(party)

				return nil
			},
		}

		distCmd.Flags().IntVarP(&gp, "gold", "g", 0, "Amount of gold")
		distCmd.Flags().IntVarP(&cp, "copper", "c", 0, "Amount of copper")
		distCmd.Flags().IntVarP(&sp, "silver", "s", 0, "Amount of silver")
		distCmd.Flags().IntVarP(&ep, "electrum", "e", 0, "Amount of electrum")
		distCmd.Flags().IntVarP(&pp, "platinum", "p", 0, "Amount of platinum")

		distCmd.Flags().BoolVarP(&equal, "force-equal", "E", false, "Do not distribute remainders")

		distCmd.Flags().StringVarP(&fn, "file", "f", "party.toml", "File to read party data from")
	}

	return distCmd
}

func Distribute(party *purse.Party) error {
	total := calc.ConvertToCopper(sp, ep, gp, pp) + cp

	extra := party.Split(total)

	if !equal {
		party.Unclaimed += extra
	} else {
		for _, p := range party.Players {
			if extra == 0 {
				break
			}

			p.Value++
			extra--
		}
	}

	return nil
}
