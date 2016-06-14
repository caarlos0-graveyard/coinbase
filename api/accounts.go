package api

import (
	"encoding/json"
	"errors"
)

type getAccountsResponse struct {
	Data []Account `json:"data"`
	Errors
}

// Account JSON
type Account struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Primary       bool   `json:"primary"`
	Currency      string `json:"currency"`
	Balance       Money  `json:"balance"`
	NativeBalance Money  `json:"native_balance"`
}

// ErrAccountNotFound is thrown when a acc with the given id was not found
var ErrAccountNotFound = errors.New("account not found")

// EmptyAccount result
var EmptyAccount = Account{}

// EmptyAccounts result
var EmptyAccounts = []Account{}

// Accounts get all accs
func (c *Client) Accounts() ([]Account, error) {
	res, err := c.Get("/accounts")
	if err != nil {
		return EmptyAccounts, err
	}
	defer res.Body.Close()
	var result getAccountsResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if res.StatusCode == 200 {
		return result.Data, err
	}
	return EmptyAccounts, c.newAPIError(res.Status, result.Errors)
}

func (c *Client) findAccount(id string) (Account, error) {
	accs, err := c.Accounts()
	if err != nil {
		return EmptyAccount, err
	}
	if id == "" {
		return accs[0], nil
	}
	for _, acc := range accs {
		if acc.ID == id {
			return acc, nil
		}
	}
	return EmptyAccount, ErrAccountNotFound
}
