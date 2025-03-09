package cmd

import (
	"github.com/asphaltbuffet/dndnd/pkg/purse"
	"github.com/spf13/cobra"
)

var createCommand *cobra.Command

var players []string

func GetCreateCommand() *cobra.Command {
	if createCommand == nil {
		createCommand = &cobra.Command{
			Use:          "create",
			Short:        "define a new party",
			Args:         cobra.ExactArgs(1),
			Aliases:      []string{"c", "new", "n"},
			SilenceUsage: true,
			RunE:         CreatePurses,
		}

		createCommand.Flags().StringSliceVarP(&players, "player", "p", []string{}, "party members' names")
		_ = createCommand.MarkFlagRequired("player")
	}

	return createCommand
}

func CreatePurses(cmd *cobra.Command, args []string) error {
	pp := &purse.Party{
		Name:    args[0],
		Players: make(map[string]purse.Purse, len(players)),
	}

	for _, name := range players {
		p := purse.NewPurse(name)

		pp.Players[name] = *p
	}

	cmd.Println(pp)

	return nil
}
