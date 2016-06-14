package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

// TransactionList cmd
var TransactionList = cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "list your last transactions",
	Flags: []cli.Flag{
		nativeFlag,
		cli.StringFlag{
			Name:  "acc, id",
			Usage: "account id (you can get it from balance)",
		},
	},
	Action: func(c *cli.Context) error {
		native := c.Bool("native")
		transactions, err := client(c).Transactions(c.String("acc"))
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		for _, transaction := range transactions {
			status := "?"
			if transaction.Status == "completed" {
				status = "✓"
			}
			fmt.Printf("%s\t", status)
			printfMoney(native, transaction.Amount, transaction.NativeAmount)
			fmt.Printf("%s\t", transaction.Created.Format("2006-01-02 15:04"))
			fmt.Printf("%s\n", transaction.Description)
		}
		return nil
	},
}
