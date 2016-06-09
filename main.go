package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/coinbase/api"
	"github.com/urfave/cli"
)

func coinbase(key, secret string) *api.Client {
	client, err := api.New(key, secret)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

const version = "dev"

func main() {
	var key string
	var secret string
	app := cli.NewApp()
	app.Name = "coinbase"
	app.Version = version
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
				if native {
					balance, err = coinbase(key, secret).NativeBalance()
				} else {
					balance, err = coinbase(key, secret).Balance()
				}
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
