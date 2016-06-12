package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

// Balance cmd
var Balance = cli.Command{
	Name:    "balance",
	Aliases: []string{"bal", "b"},
	Usage:   "get your balance",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "native, n",
			Usage: "Native Currency balance",
		},
		cli.BoolFlag{
			Name:  "id",
			Usage: "Show account id",
		},
	},
	Action: func(c *cli.Context) error {
		native := c.Bool("native")
		showID := c.Bool("id")
		accounts, err := coinbase(c).Accounts()
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		for _, acc := range accounts {
			if showID {
				fmt.Printf("%s\t", acc.ID)
			}
			fmt.Printf("%s\t", acc.Name)
			BalancePrintf(native, acc)
			fmt.Printf("\n")
		}
		return nil
	},
}
