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
		assert.NotNil(t, transactions)
	}
}

func TestTransferMoney(t *testing.T) {
	cli := NewTestCli()
	accs, _ := cli.Accounts()
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
	assert.Equal(t, tx.Status, "completed")
}

func TestSendMoney(t *testing.T) {
	cli := NewTestCli()
	accs, _ := cli.Accounts()
	var from api.Account
	for _, acc := range accs {
		if acc.Name == "BTC Wallet" {
			from = acc
			break
		}
	}
	tx, err := cli.Send(
		from.ID,
		"n4VQ5YdHf7hLQ2gWQYYrcxoE5B7nWuDFNF",
		"0.010232",
		"BTC",
		"more money",
	)
	assert.NoError(t, err)
	assert.Equal(t, tx.Status, "pending")
}

func TestSendMoneyInvalidAddr(t *testing.T) {
	cli := NewTestCli()
	accs, _ := cli.Accounts()
	var from api.Account
	for _, acc := range accs {
		if acc.Name == "BTC Wallet" {
			from = acc
			break
		}
	}
	tx, err := cli.Send(
		from.ID,
		"1dce72d353d64f81833f8cd318405766cb3f09c828fb2e85edf41ed05aded467",
		"0.010232",
		"BTC",
		"more money",
	)
	assert.Error(t, err)
	assert.Equal(t, tx, api.Transaction{})
}
