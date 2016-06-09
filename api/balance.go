package api

import (
	"encoding/json"
	"net/http"
)

type Balance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type accounts struct {
	Data []struct {
		Balance       Balance `json:"balance"`
		NativeBalance Balance `json:"native_balance"`
	} `json:"data"`
}

func (c *Client) Balance() (Balance, error) {
	accs, err := c.getAccounts()
	return accs.Data[0].Balance, err
}

func (c *Client) NativeBalance() (Balance, error) {
	accs, err := c.getAccounts()
	return accs.Data[0].NativeBalance, err
}

func (c *Client) getAccounts() (accounts, error) {
	req, err := http.NewRequest(
		"GET",
		c.BaseURL+"/accounts",
		nil,
	)
	if err != nil {
		return accounts{}, err
	}
	res, err := c.Do(req)
	if err != nil {
		return accounts{}, err
	}
	defer res.Body.Close()
	var result accounts
	return result, json.NewDecoder(res.Body).Decode(&result)
}
