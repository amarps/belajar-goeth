package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKs() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "./UTC--2024-03-27T09-21-18.526155000Z--899ceb42957102ee697c79990dc88aa5eb45fdfa"
	ks := keystore.NewKeyStore("./wallets/", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		panic(err)
	}
}

func main() {
	importKs()
}