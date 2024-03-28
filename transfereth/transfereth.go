package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"goeth/setup"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)


func main() {
	setup.Init()
    client, err := ethclient.Dial(setup.EnvEthNet())
    if err != nil {
        panic(err)
    }

    privateKey, err := crypto.HexToECDSA("bf51ebc853cd1619351dea7147e3f1034809e00a16575224cf2bcef213866d39")
    if err != nil {
        panic(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        panic(err)
    }

    value := big.NewInt(1000000000000000000) // in wei (1 eth)
    gasLimit := uint64(21000)                // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        panic(err)
    }

    toAddress := common.HexToAddress("0xa44a982026875f838Be23092A2F50fe46E77EEbc")
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

    chainID, err := client.ChainID(context.Background())
    if err != nil {
        panic(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        panic(err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        panic(err)
    }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}