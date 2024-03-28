package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

	fmt.Println("we have a connection")
	_ = client
}