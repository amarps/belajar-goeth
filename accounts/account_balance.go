package main

import (
	"context"
	"fmt"
	"goeth/setup"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func main() {
	setup.Init()

	ethAddr := viper.GetString("ETH_NET")
	client, err := ethclient.Dial(ethAddr)
	if err != nil {
		panic(err)
	}

	account := common.HexToAddress("0x6dCfBE6eFa1D49F83E338453Fd61f9c548dA7CdD")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		panic(err)
	}
	fmt.Println(pendingBalance)
}