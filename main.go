package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/coinbase/api"
	"github.com/fabioberger/coinbase-go"
	"github.com/urfave/cli"
)

func CoinbaseClient(key, secret string) coinbase.Client {
	fmt.Println(key)
	fmt.Println(secret)
	return coinbase.ApiKeyClient(key, secret)
}

func main() {
	var key string
	var secret string
	app := cli.NewApp()
	app.Name = "coinbase"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "key",
			Usage:       "Coinbase API Key",
			EnvVar:      "COINBASE_API_KEY",
			Destination: &key,
		},
		cli.StringFlag{
			Name:        "secret",
			Usage:       "Coinbase API Secret",
			EnvVar:      "COINBASE_API_SECRET",
			Destination: &secret,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "balance",
			Aliases: []string{"b"},
			Usage:   "get your balance",
			Action: func(c *cli.Context) error {
				cc, err := api.New(key, secret)
				if err != nil {
					fmt.Println(err)
				}
				balance, err := cc.Balance()
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				fmt.Printf("%s %s", balance.Amount, balance.Currency)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
