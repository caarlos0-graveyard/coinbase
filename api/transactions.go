package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type transactions struct {
	Data []Transaction `json:"data"`
}

// Transaction JSON
type Transaction struct {
	ID           string    `json:"id"`
	Type         string    `json:"type"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	Amount       Money     `json:"amount"`
	NativeAmount Money     `json:"native_amount"`
	Created      time.Time `json:"created_at"`
	Updated      time.Time `json:"updated_at"`
}

// Transactions for a given account
func (c *Client) Transactions(accountID string) ([]Transaction, error) {
	acc, err := c.account(accountID)
	if err != nil {
		return []Transaction{}, err
	}
	req, err := http.NewRequest(
		"GET", c.BaseURL+"/accounts/"+acc.ID+"/transactions", nil,
	)
	if err != nil {
		return []Transaction{}, err
	}
	res, err := c.Do(req)
	if err != nil {
		return []Transaction{}, err
	}
	defer res.Body.Close()
	var result transactions
	return result.Data, json.NewDecoder(res.Body).Decode(&result)
}
