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

func (c *Client) Epoch() (string, error) {
	var result apiTime
	resp, err := c.client.Get(c.BaseURL + "/time")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return fmt.Sprintf("%d", result.Data.Epoch), err
}
