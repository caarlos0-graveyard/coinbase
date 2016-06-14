package api_test

import (
	"testing"

	"github.com/caarlos0/coinbase/api"
	"github.com/stretchr/testify/assert"
)

func TestListAccounts(t *testing.T) {
	accs, err := NewTestCli().Accounts()
	assert.NoError(t, err)
	assert.NotEmpty(t, accs)
}

func TestListAccountsError(t *testing.T) {
	cli, err := api.New("wrong", "also-wrong")
	assert.Nil(t, err)
	accs, err := cli.Accounts()
	assert.Error(t, err)
	assert.Empty(t, accs)
}
