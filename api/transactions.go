package api

import (
	"encoding/json"
	"errors"
	"time"
)

type transactions struct {
	Data []Transaction `json:"data"`
}

type transaction struct {
	Data Transaction `json:"data"`
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

// Request JSON
type Request struct {
	Type        string `json:"type"`
	To          string `json:"to"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Idem        string `json:"idem"`
}

// Transactions for a given account
func (c *Client) Transactions(accountID string) ([]Transaction, error) {
	acc, err := c.findAccount(accountID)
	if err != nil {
		return []Transaction{}, err
	}
	res, err := c.Get("/accounts/" + acc.ID + "/transactions")
	if err != nil {
		return []Transaction{}, err
	}
	defer res.Body.Close()
	var result transactions
	return result.Data, json.NewDecoder(res.Body).Decode(&result)
}

// Transfer money from an acc to another
func (c *Client) Transfer(
	from, to, amount, currency, description string,
) (Transaction, error) {
	acc, err := c.findAccount(from)
	if err != nil {
		return Transaction{}, err
	}
	request := Request{
		Type:        "transfer",
		To:          to,
		Amount:      amount,
		Currency:    currency,
		Description: description,
	}
	res, err := c.Post("/accounts/"+acc.ID+"/transactions", request)
	if err != nil {
		return Transaction{}, err
	}
	defer res.Body.Close()
	if res.StatusCode == 201 {
		var result transaction
		return result.Data, json.NewDecoder(res.Body).Decode(&result)
	}

	return Transaction{}, errors.New(res.Status)
}
