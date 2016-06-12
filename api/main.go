package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"html"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// Client to coinbase api
type Client struct {
	client               *http.Client
	key, secret, baseURL string
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
		baseURL: baseURL(),
	}, nil
}

func baseURL() string {
	if os.Getenv("COINBASE_SANDBOX") != "" {
		log.Println("Using sandbox environment...")
		return "https://api.sandbox.coinbase.com/v2"
	}
	return "https://api.coinbase.com/v2"
}

// UnsignedGet a path, not signed
func (c *Client) UnsignedGet(path string) (*http.Response, error) {
	return c.client.Get(c.baseURL + path)
}

// Get a path, signed
func (c *Client) Get(path string) (*http.Response, error) {
	timestamp, err := c.Epoch()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	c.appendHeaders(req, timestamp, "")
	return c.client.Do(req)
}

func (c *Client) appendHeaders(req *http.Request, timestamp, body string) {
	req.Header.Set("User-Agent", "BeckerCoinbase/v1")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("CB-VERSION", "2016-06-11")
	req.Header.Set("CB-ACCESS-KEY", c.key)
	req.Header.Set("CB-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("CB-ACCESS-SIGN", c.sign(timestamp, req.Method, html.EscapeString(req.URL.Path), ""))
}

func (c *Client) sign(timestamp, method, path, body string) string {
	h := hmac.New(sha256.New, []byte(c.secret))
	h.Write([]byte(timestamp + method + path + body))
	return hex.EncodeToString(h.Sum(nil))
}

func dialTimeout(network, addr string) (net.Conn, error) {
	var timeout = time.Duration(20 * time.Second)
	return net.DialTimeout(network, addr, timeout)
}
