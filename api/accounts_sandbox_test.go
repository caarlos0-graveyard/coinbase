package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAccounts(t *testing.T) {
	accs, err := NewTestCli().Accounts()
	assert.Nil(t, err)
	assert.NotEmpty(t, accs)
}
