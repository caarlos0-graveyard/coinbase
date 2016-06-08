package main

import (
	"fmt"
	"os"

	"github.com/fabioberger/coinbase-go"
	"github.com/urfave/cli"
)

func CoinbaseClient(key, token string) coinbase.Client {
	fmt.Println(key)
	fmt.Println(token)
	return coinbase.ApiKeyClient(key, token)
}

func main() {
	var key string
	var token string
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
			Name:        "token",
			Usage:       "Coinbase API Token",
			EnvVar:      "COINBASE_API_TOKEN",
			Destination: &token,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "balance",
			Aliases: []string{"b"},
			Usage:   "get your balance",
			Action: func(c *cli.Context) error {
				bal, err := CoinbaseClient(key, token).GetBalance()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("%f BTC", bal)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
