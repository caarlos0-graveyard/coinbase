package api

import (
	"encoding/json"
	"fmt"
)

type apiTime struct {
	Data struct {
		Epoch int `json:"epoch"`
	} `json:"data"`
}

// Epoch of the server
func (c *Client) Epoch() (string, error) {
	var result apiTime
	resp, err := c.UnsignedGet("/time")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return fmt.Sprintf("%d", result.Data.Epoch), err
}
