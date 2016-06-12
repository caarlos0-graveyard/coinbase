package api_test

import (
	"log"
	"os"

	"github.com/caarlos0/coinbase/api"
)

func NewTestCli() *api.Client {
	cli, err := api.New(
		os.Getenv("COINBASE_SANDBOX_KEY"),
		os.Getenv("COINBASE_SANDBOX_SECRET"),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return cli
}
