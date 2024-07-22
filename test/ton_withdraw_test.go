package jetton

import (
	"fmt"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"log"
	"testing"
	"ton-demo/utils"
)

func TestDeployTonWithdraw(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())

	msgBody := cell.BeginCell().EndCell()

	fmt.Println("Deploying ton withdraw contract to testnet...")
	tonWithdrawBoc := "b5ee9c724101060100a5000114ff00f4a413f4bcf2c80b01020162020300e6d032d0d30331fa403001d31fed44d0d31ffa40fa40303123c0018e113331a45902c8cb1f01cf1601cf16c9ed54e03122c002925f04e002c0038e325222c705f2e067fa0030f8276f22305301bef2e0688208989680a1b60871708018c8cb055004cf1658fa0212cb6ac901fb00e05f03f2c3090201580405000fb96c0f8276f223080019bba4fed44d0d31ffa40fa40308c646791e"
	addr, _, _, err := w.DeployContractWaitTransaction(ctx, tlb.MustFromTON("0.3"),
		msgBody, getJettonCode(tonWithdrawBoc), getTonWithdrawContractData(w.WalletAddress()))
	if err != nil {
		panic(err)
	}

	fmt.Println("Deployed contract addr:", addr.String())
}

func TestCounter(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())
	contractAddr := address.MustParseAddr("EQAz22xirtU0TxXusSSzYW3SHuidr08FgfcbqQosfTcc1g4I")
	fmt.Println("deposit ton...")
	depositPayload := cell.BeginCell().MustStoreUInt(1, 32).EndCell()
	counter := wallet.SimpleMessageAutoBounce(contractAddr, tlb.MustFromTON("0.001"), depositPayload)
	_, _, err := w.SendWaitTransaction(ctx, counter)
	if err != nil {
		panic(err)
	}
}

func TestDeposit(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())
	contractAddr := address.MustParseAddr("EQAz22xirtU0TxXusSSzYW3SHuidr08FgfcbqQosfTcc1g4I")
	fmt.Println("deposit ton...")
	depositPayload := cell.BeginCell().MustStoreUInt(2, 32).EndCell()
	deposit := wallet.SimpleMessageAutoBounce(contractAddr, tlb.MustFromTON("0.03"), depositPayload)
	_, _, err := w.SendWaitTransaction(ctx, deposit)
	if err != nil {
		panic(err)
	}
}

func TestWithdraw(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())
	contractAddr := address.MustParseAddr("EQAz22xirtU0TxXusSSzYW3SHuidr08FgfcbqQosfTcc1g4I")

	fmt.Println("withdraw ton...")
	withdrawPayload := cell.BeginCell().MustStoreUInt(3, 32).
		MustStoreCoins(tlb.MustFromTON("0.12").Nano().Uint64()).EndCell()
	withdraw := wallet.SimpleMessageAutoBounce(contractAddr, tlb.MustFromTON("0.01"), withdrawPayload)
	_, _, err := w.SendWaitTransaction(ctx, withdraw)
	if err != nil {
		panic(err)
	}
}

func TestGetBalance(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("wallet:", w.WalletAddress().String())
	contractAddr := address.MustParseAddr("EQAz22xirtU0TxXusSSzYW3SHuidr08FgfcbqQosfTcc1g4I")
	// we need fresh block info to run get methods
	b, err := api.CurrentMasterchainInfo(ctx)
	if err != nil {
		log.Fatalln("get block err:", err.Error())
		return
	}
	res, err := api.WaitForBlock(b.SeqNo).RunGetMethod(ctx, b,
		contractAddr, "balance")
	if err != nil {
		log.Fatalln("run get method err:", err.Error())
		return
	}

	balance, err := res.Int(0)
	if err != nil {
		println("ERR", err.Error())
		return
	}
	fmt.Printf("balance result: %d\n", balance)
}

func TestGetContractStorageData(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("wallet:", w.WalletAddress().String())
	contractAddr := address.MustParseAddr("EQAz22xirtU0TxXusSSzYW3SHuidr08FgfcbqQosfTcc1g4I")
	// we need fresh block info to run get methods
	b, err := api.CurrentMasterchainInfo(ctx)
	if err != nil {
		log.Fatalln("get block err:", err.Error())
		return
	}
	res, err := api.WaitForBlock(b.SeqNo).RunGetMethod(ctx, b,
		contractAddr, "get_contract_storage_data")
	if err != nil {
		log.Fatalln("run get method err:", err.Error())
		return
	}

	counter, err := res.Int(0)
	if err != nil {
		println("ERR", err.Error())
		return
	}
	fmt.Printf("counter result: %d\n", counter)
}

func getTonWithdrawContractData(ownerAddr *address.Address) *cell.Cell {
	data := cell.BeginCell().MustStoreUInt(0, 32).MustStoreAddr(ownerAddr).MustStoreAddr(ownerAddr).EndCell()
	return data
}
