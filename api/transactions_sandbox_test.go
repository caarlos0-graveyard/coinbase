package api_test

import (
	"log"
	"testing"

	"github.com/caarlos0/coinbase/api"
	"github.com/stretchr/testify/assert"
)

func TestListTransactions(t *testing.T) {
	cli := NewTestCli()
	accs, err := cli.Accounts()
	assert.Nil(t, err)
	for _, acc := range accs {
		log.Println("Getting transactions for", acc.ID)
		transactions, err := cli.Transactions(acc.ID)
		assert.Nil(t, err)
		log.Println(transactions)
	}
}

func TestTransferMoney(t *testing.T) {
	cli := NewTestCli()
	accs, err := cli.Accounts()
	assert.Nil(t, err)
	var from api.Account
	var to api.Account
	for _, acc := range accs {
		if acc.Name == "BTC Wallet" {
			from = acc
		}
		if acc.Name == "BTC Wallet 1" {
			to = acc
		}
	}
	tx, err := cli.Transfer(from.ID, to.ID, "0.001", "BTC", "some money")
	assert.Nil(t, err)
	log.Println(tx)
}
