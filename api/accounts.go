package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type accounts struct {
	Data []Account `json:"data"`
}

// Account JSON
type Account struct {
	ID            string `json:"id"`
	Balance       Money  `json:"balance"`
	NativeBalance Money  `json:"native_balance"`
}

// ErrAccountNotFound is thrown when a acc with the given id was not found
var ErrAccountNotFound = errors.New("account not found")

// Accounts get all accs
func (c *Client) Accounts() ([]Account, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/accounts", nil)
	if err != nil {
		return []Account{}, err
	}
	res, err := c.Do(req)
	if err != nil {
		return []Account{}, err
	}
	defer res.Body.Close()
	var result accounts
	return result.Data, json.NewDecoder(res.Body).Decode(&result)
}

func (c *Client) account(id string) (Account, error) {
	accs, err := c.Accounts()
	if err != nil {
		return Account{}, err
	}
	if id == "" {
		return accs[0], nil
	}
	for _, acc := range accs {
		if acc.ID == id {
			return acc, nil
		}
	}
	return Account{}, ErrAccountNotFound
}
