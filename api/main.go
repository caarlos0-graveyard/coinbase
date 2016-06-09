package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"html"
	"net"
	"net/http"
	"time"
)

// Client to coinbase api
type Client struct {
	client               *http.Client
	key, secret, BaseURL string
}

// New secure client
func New(key, secret string) (*Client, error) {
	pool := x509.NewCertPool()
	if ok := pool.AppendCertsFromPEM(certs); !ok {
		return nil, errors.New("failed to parse certs")
	}
	return &Client{
		client: &http.Client{
			Transport: &http.Transport{
				Dial:            dialTimeout,
				TLSClientConfig: &tls.Config{RootCAs: pool},
			},
		},
		key:     key,
		secret:  secret,
		BaseURL: "https://api.coinbase.com/v2",
	}, nil
}

// Do a signed request
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	timestamp, err := c.Epoch()
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "BeckerCoinbase/v1")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("CB-ACCESS-KEY", c.key)
	req.Header.Set("CB-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("CB-ACCESS-SIGN", c.sign(timestamp, req))

	return c.client.Do(req)
}

func (c *Client) sign(timestamp string, req *http.Request) string {
	// TODO missing body contents on the mix
	message := timestamp + req.Method + html.EscapeString(req.URL.Path) + ""
	h := hmac.New(sha256.New, []byte(c.secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func dialTimeout(network, addr string) (net.Conn, error) {
	var timeout = time.Duration(20 * time.Second)
	return net.DialTimeout(network, addr, timeout)
}
