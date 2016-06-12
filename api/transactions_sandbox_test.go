package api_test

import (
	"log"
	"testing"

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
