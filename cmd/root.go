package cmd

import (
	"os"
	"strconv"

	"github.com/asphaltbuffet/dndnd/pkg/purse"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func Execute() {
	if err := GetRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func GetRootCommand() *cobra.Command {
	if rootCmd == nil {
		rootCmd = &cobra.Command{
			Use:          "dndnd",
			Short:        "wealth distribution for the party",
			SilenceUsage: true,
			Version:      "0.0.1",
			Args:         cobra.NoArgs,
			RunE: func(cmd *cobra.Command, _ []string) error {
				pp, err := purse.Load(fn)
				if err != nil {
					return err
				}

				cmd.Println(pp)

				return nil
			},
		}

		rootCmd.AddCommand(GetCreateCommand())
		rootCmd.AddCommand(GetDistributeCommand())
	}

	return rootCmd
}

func printPurses(purses []*purse.Purse) string {
	data := [][]string{
		{"", "PP", "GP", "EP", "SP", "CP"},
	}

	for i, p := range purses {
		r := []string{
			"Player " + strconv.Itoa(i+1),
			strconv.Itoa(p.Platinum()),
			strconv.Itoa(p.Gold()),
			strconv.Itoa(p.Electrum()),
			strconv.Itoa(p.Silver()),
			strconv.Itoa(p.Copper()),
		}
		data = append(data, r)
	}

	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1).Align(lipgloss.Right)
	headerStyle := baseStyle.Foreground(lipgloss.Color("247")).Bold(true)

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return baseStyle.Foreground(lipgloss.Color("92"))
			default:
				return baseStyle.Foreground(lipgloss.Color("46"))
			}
		}).
		Headers(data[0]...).
		Rows(data[1:]...)

	return t.Render()
}
