package CryptoPrice

import "fmt"

type Asset struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	PriceUSD string `json:"priceUsd"`
}

func (d Asset) GetInfo() string {
	return fmt.Sprintf("%s (%s) price: $%s", d.Name, d.Symbol, d.PriceUSD)
}

type assetResponse struct {
	Data      Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
