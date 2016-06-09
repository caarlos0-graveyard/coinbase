package commands

import (
	"log"

	"github.com/caarlos0/coinbase/api"
	"github.com/urfave/cli"
)

func coinbase(c *cli.Context) *api.Client {
	client, err := api.New(
		c.Parent().String("key"),
		c.Parent().String("secret"),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
