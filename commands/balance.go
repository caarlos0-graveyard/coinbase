package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

var Balance = cli.Command{
	Name:    "balance",
	Aliases: []string{"bal", "b"},
	Usage:   "get your balance",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "native, n",
			Usage: "Native Currency balance",
		},
	},
	Action: func(c *cli.Context) error {
		native := c.Bool("native")
		accounts, err := coinbase(c).Accounts()
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		for _, acc := range accounts {
			// FIXME improve this
			if native {
				fmt.Printf(
					"%s\t%s\t%s\n",
					acc.ID,
					acc.NativeBalance.Amount,
					acc.NativeBalance.Currency,
				)
			} else {
				fmt.Printf(
					"%s\t%s\t%s\n",
					acc.ID,
					acc.Balance.Amount,
					acc.Balance.Currency,
				)
			}
		}
		return nil
	},
}
