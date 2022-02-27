package main

import (
	"api-client/CryptoPrice"
	"fmt"
	"log"
	"time"
)

func main() {
	apiClient, err := CryptoPrice.NewClient(time.Second * 3)
	if err != nil {
		log.Fatal(err)
	}

	var coinName string
	fmt.Print("Enter the coin's name: ")
	fmt.Scan(&coinName)
	coin, err := apiClient.GetAsset(coinName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coin.GetInfo())
}
