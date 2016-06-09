package commands

import (
	"fmt"

	"github.com/caarlos0/coinbase/api"
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
		var balance api.Balance
		var err error
		client := coinbase(c)
		if native {
			balance, err = client.NativeBalance()
		} else {
			balance, err = client.Balance()
		}
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		fmt.Printf("%s %s", balance.Amount, balance.Currency)
		return nil
	},
}
