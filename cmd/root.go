package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	gold     int
	copper   int
	silver   int
	electrum int
	platinum int
)

var rootCmd *cobra.Command

func Execute() {
	if err := GetRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetRootCommand() *cobra.Command {
	if rootCmd == nil {
		rootCmd = &cobra.Command{
			Use:     "dndnd",
			Short:   "A brief description of your application",
			Version: "0.0.1",
			Args:    cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Gold: %d, Copper: %d, Silver: %d, Electrum: %d, Platinum: %d\n", gold, copper, silver, electrum, platinum)
			},
		}

		rootCmd.Flags().IntVarP(&gold, "gold", "g", 0, "Amount of gold")
		rootCmd.Flags().IntVarP(&copper, "copper", "c", 0, "Amount of copper")
		rootCmd.Flags().IntVarP(&silver, "silver", "s", 0, "Amount of silver")
		rootCmd.Flags().IntVarP(&electrum, "electrum", "e", 0, "Amount of electrum")
		rootCmd.Flags().IntVarP(&platinum, "platinum", "p", 0, "Amount of platinum")

	}

	return rootCmd
}
