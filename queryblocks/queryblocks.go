package main

import (
	"context"
	"fmt"
	"goeth/setup"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	setup.Init()
	client, err := ethclient.Dial(setup.EnvEthNet())
	if err != nil {
		panic(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(header.Number.String())
}