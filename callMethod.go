package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	test "github.com/AleksaOpacic/sendTx/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://localhost:10002")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("4a3a1205239c37f52fa260e4c2878b62de55a8104ac2d012165013922d6c1b9f")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x563e8b35Cd4ED52a6d9c2c97D1E5E104a5d39F8e")
	instance, err := test.NewTest(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	copy(key[:], []byte("0x6261720000000000000000000000000000000000000000000000000000000000"))

	tx, err := instance.Deposit(auth, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx sent: ", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

}
