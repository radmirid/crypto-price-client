package CryptoPrice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("error: zero timeout")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &logRT{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetAsset(name string) (Asset, error) {
	url := fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)

	resp, err := c.client.Get(url)
	if err != nil {
		return Asset{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Asset{}, err
	}

	var r assetResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return Asset{}, err
	}

	return r.Data, nil
}
