#  Crypto Price Client

Golang REST API client which outputs the price of the entered cryptocurrency.

## Installing

```
go get github.com/radmirid/crypto-price-client
```

## Running

```
go run main.go
```

##  Usage Example

`Input`

```
Enter the coin's name: solana
```

`Output`

```
Sun Feb 27 12:00:00 2022 | GET | https://api.coincap.io/v2/assets/solana
Solana (SOL) price: $89.2912533674401356
```

## LICENSE

[MIT License](LICENSE)
