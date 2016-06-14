package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// TransactionSend cmd
var TransactionSend = cli.Command{
	Name:    "send",
	Aliases: []string{"mv"},
	Usage:   "Send money to a BTC address",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "acc, id",
			Usage: "account id (you can get it from balance)",
		},
		cli.StringFlag{
			Name:  "to",
			Usage: "BTC address to send money to",
		},
		cli.StringFlag{
			Name:  "amount",
			Usage: "Amount of money to send",
		},
		cli.StringFlag{
			Name:  "currency",
			Usage: "Currency of money to send (e.g. BTC, USD, BRL)",
			Value: "BTC",
		},
		cli.StringFlag{
			Name:  "description",
			Usage: "Description of the transaction",
		},
	},
	Action: func(c *cli.Context) error {
		from := c.String("acc")
		to := c.String("to")
		amount := c.String("amount")
		currency := c.String("currency")
		description := c.String("description")
		fmt.Printf("Sending %s %s to %s...\n", amount, currency, to)
		fmt.Printf("Press 'Enter' to confirm or 'CTRL-C' to cancel...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		transaction, err := client(c).Send(
			from,
			to,
			amount,
			currency,
			description,
		)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		fmt.Println(transaction)
		return nil
	},
}
