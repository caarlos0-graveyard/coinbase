package main

import (
	"os"

	"github.com/caarlos0/coinbase/commands"
	"github.com/urfave/cli"
)

const version = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "coinbase"
	app.Author = "Carlos Alexandro Becker (@caarlos0)"
	app.Usage = "Client for the Coinbase API"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "key",
			Usage:  "Coinbase API Key",
			EnvVar: "COINBASE_API_KEY",
		},
		cli.StringFlag{
			Name:   "secret",
			Usage:  "Coinbase API Secret",
			EnvVar: "COINBASE_API_SECRET",
		},
	}
	app.Commands = []cli.Command{
		commands.Balance,
		commands.TransactionList,
		commands.TransactionSend,
	}
	app.Run(os.Args)
}
