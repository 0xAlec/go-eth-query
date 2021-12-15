package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	token "main/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

const fwb = "0x35bd01fc9d6d5d81ca9e055db88dc49aa2c699a8"
const fiftyk = "0x15d09449fadd279d65a43c49980cd3992c92cb49"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	endpoint := os.Getenv("API")

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	tokenAddr := common.HexToAddress(fwb)
	instance, err := token.NewToken(tokenAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	userAddr := common.HexToAddress(fiftyk)

	bal, err := instance.BalanceOf(&bind.CallOpts{}, userAddr)
	if err != nil {
		log.Fatal(err)
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
	fmt.Printf("Account %s Balance: %f tokens\n", userAddr, value)
}
