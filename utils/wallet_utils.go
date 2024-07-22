package utils

import (
	"context"
	"encoding/hex"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"strings"
)

func GetWallet(api ton.APIClientWrapped) *wallet.Wallet {
	words := strings.Split("erosion sense trend impose bulb receive eye program access divert opera sea liquid bounce render street hobby enhance afford engage clean dirt now six", " ")
	//words = strings.Split("sad hobby cruel quality delay story install mixture prefer opera clip credit wrap rally defense regular swift hero town swing bottom blossom flush globe", " ")
	w, err := wallet.FromSeed(api, words, wallet.V4R2)
	if err != nil {
		panic(err)
	}
	return w
}

func GetCodeCell(hexBOC string) *cell.Cell {
	codeCellBytes, _ := hex.DecodeString(hexBOC)
	codeCell, err := cell.FromBOC(codeCellBytes)
	if err != nil {
		panic(err)
	}
	return codeCell
}

func GetTestApiAndClient(isTest bool) (*ton.APIClient, context.Context) {
	client := liteclient.NewConnectionPool()

	// connect to testnet lite server
	mainConfigUrl := "https://ton.org/global.config.json"
	testConfigUrl := "https://ton-blockchain.github.io/testnet-global.config.json"
	if isTest {
		err := client.AddConnectionsFromConfigUrl(context.Background(), testConfigUrl)
		if err != nil {
			panic(err)
		}
	} else {
		err := client.AddConnectionsFromConfigUrl(context.Background(), mainConfigUrl)
		if err != nil {
			panic(err)
		}
	}
	ctx := client.StickyContext(context.Background())
	// initialize ton api lite connection wrapper with full proof checks
	api := ton.NewAPIClient(client)
	return api, ctx
}
