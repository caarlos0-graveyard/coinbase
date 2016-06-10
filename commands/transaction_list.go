package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

var TransactionList = cli.Command{
	Name:    "transaction-list",
	Aliases: []string{"list", "ls", "l"},
	Usage:   "list your last transactions",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "native, n",
			Usage: "Native Currency balance",
		},
		cli.StringFlag{
			Name:  "account-id, acc, id",
			Usage: "account id (you can get it from balance)",
		},
	},
	Action: func(c *cli.Context) error {
		native := c.Bool("native")
		transactions, err := coinbase(c).Transactions(c.String("accounts-id"))
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		for _, transaction := range transactions {
			status := "?"
			if transaction.Status == "completed" {
				status = "✓"
			}
			fmt.Printf("%s\t", status)
			MoneyPrintf(native, transaction.NativeAmount, transaction.Amount)
			fmt.Printf("%s\t", transaction.Created.Format("2006-01-02 15:04"))
			fmt.Printf("%s\n", transaction.Description)
		}
		return nil
	},
}
