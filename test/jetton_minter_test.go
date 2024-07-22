package jetton

import (
	"encoding/hex"
	"fmt"
	"github.com/e9571/lib1"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton/jetton"
	"github.com/xssnick/tonutils-go/ton/nft"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"log"
	"testing"
	"ton-demo/utils"
)

func TestDeployJettonMint(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())

	msgBody := cell.BeginCell().EndCell()

	fmt.Println("Deploying jetton contract to testnet...")
	jettonWalletBoc := "b5ee9c7241021201000334000114ff00f4a413f4bcf2c80b0102016202110202cc03060201d4040500c30831c02497c138007434c0c05c6c2544d7c0fc02f83e903e900c7e800c5c75c87e800c7e800c1cea6d0000b4c7e08403e29fa954882ea54c4d167c0238208405e3514654882ea58c511100fc02780d60841657c1ef2ea4d67c02b817c12103fcbc2000113e910c1c2ebcb85360020148070e020120080a01f100f4cffe803e90087c007b51343e803e903e90350c144da8548ab1c17cb8b04a30bffcb8b0950d109c150804d50500f214013e809633c58073c5b33248b232c044bd003d0032c032483e401c1d3232c0b281f2fff274013e903d010c7e800835d270803cb8b11de0063232c1540233c59c3e8085f2dac4f3200900ae8210178d4519c8cb1f19cb3f5007fa0222cf165006cf1625fa025003cf16c95005cc2391729171e25008a813a08208e4e1c0aa008208989680a0a014bcf2e2c504c98040fb001023c85004fa0258cf1601cf16ccc9ed5403f73b51343e803e903e90350c0234cffe80145468017e903e9014d6f1c1551cdb5c150804d50500f214013e809633c58073c5b33248b232c044bd003d0032c0327e401c1d3232c0b281f2fff274140371c1472c7cb8b0c2be80146a2860822625a020822625a004ad8228608239387028062849f8c3c975c2c070c008e00b0c0d00705279a018a182107362d09cc8cb1f5230cb3f58fa025007cf165007cf16c9718010c8cb0524cf165006fa0215cb6a14ccc971fb0010241023000e10491038375f040076c200b08e218210d53276db708010c8cb055008cf165004fa0216cb6a12cb1f12cb3fc972fb0093356c21e203c85004fa0258cf1601cf16ccc9ed540201200f1000db3b51343e803e903e90350c01f4cffe803e900c145468549271c17cb8b049f0bffcb8b0a0823938702a8005a805af3cb8b0e0841ef765f7b232c7c572cfd400fe8088b3c58073c5b25c60063232c14933c59c3e80b2dab33260103ec01004f214013e809633c58073c5b3327b55200083200835c87b51343e803e903e90350c0134c7e08405e3514654882ea0841ef765f784ee84ac7cb8b174cfcc7e800c04e81408f214013e809633c58073c5b3327b5520001ba0f605da89a1f401f481f481a861f0a7d84b"
	jettonMintBoc := "b5ee9c7241020b010001ed000114ff00f4a413f4bcf2c80b0102016202080202cc030703efd9910e38048adf068698180b8d848adf07d201800e98fe99ff6a2687d007d206a6a18400aa9385d47181a9aa8aae382f9702480fd207d006a18106840306b90fd001812881a28217804502a906428027d012c678b666664f6aa7041083deecbef29385d71811a92e001f1811802600271812f82c207f978404050600fe3603fa00fa40f82854120870542013541403c85004fa0258cf1601cf16ccc922c8cb0112f400f400cb00c9f9007074c8cb02ca07cbffc9d05008c705f2e04a12a1035024c85004fa0258cf16ccccc9ed5401fa403020d70b01c3008e1f8210d53276db708010c8cb055003cf1622fa0212cb6acb1fcb3fc98042fb00915be200303515c705f2e049fa403059c85004fa0258cf16ccccc9ed54002e5143c705f2e049d43001c85004fa0258cf16ccccc9ed540093dfc142201b82a1009aa0a01e428027d012c678b00e78b666491646580897a007a00658064907c80383a6465816503e5ffe4e83bc00c646582ac678b28027d0109e5b589666664b8fd80402037a60090a007dadbcf6a2687d007d206a6a183618fc1400b82a1009aa0a01e428027d012c678b00e78b666491646580897a007a00658064fc80383a6465816503e5ffe4e840001faf16f6a2687d007d206a6a183faa9040788b22c0"
	addr, _, _, err := w.DeployContractWaitTransaction(ctx, tlb.MustFromTON("0.1"),
		msgBody, getJettonCode(jettonMintBoc), getJettonContractDataJson(w.WalletAddress(), jettonWalletBoc))
	if err != nil {
		panic(err)
	}

	fmt.Println("Deployed contract addr:", addr.String())
}

func TestMintJetton(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())
	mintAddr := address.MustParseAddr("EQDMXOy25h626He6BCy0mRM7ZvAXBQtN9U_ufStjI2GBYDnf")
	token := jetton.NewJettonMasterClient(api, mintAddr)

	jettonData, err := token.GetJettonData(ctx)
	if err != nil {
		panic(err)
	}
	utils.LogOut("info,", "jettonData:"+lib1.Json_Package(jettonData))
	toAddress := address.MustParseAddr("UQDUkb17UDGmYyQXst7oiPWdALTRlECaxgmlPioeGi7nvlPT")
	mintData, err := BuildJettonMintPayload(toAddress, w.WalletAddress(), tlb.MustFromTON("0.015"), tlb.MustFromTON("88888"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Minting jettons...")
	mint := wallet.SimpleMessageAutoBounce(mintAddr, tlb.MustFromTON("0.02"), mintData)
	_, _, err = w.SendWaitTransaction(ctx, mint)
	if err != nil {
		panic(err)
	}
}

func getJettonCode(hexBOC string) *cell.Cell {
	codeCellBytes, _ := hex.DecodeString(hexBOC)
	codeCell, err := cell.FromBOC(codeCellBytes)
	if err != nil {
		panic(err)
	}
	return codeCell
}

func getJettonContractDataJson(ownerAddr *address.Address, jettonWalletBoc string) *cell.Cell {
	jettonContent := nft.ContentOffchain{
		URI: "https://black-necessary-orca-958.mypinata.cloud/ipfs/QmNaVgXwPVda9vSXLWKK5SUQan7e4zhoNGCcrmdAewhbLH",
	}
	contentRef, _ := jettonContent.ContentCell()
	data := cell.BeginCell().MustStoreCoins(0).MustStoreAddr(ownerAddr).
		MustStoreRef(contentRef).
		MustStoreRef(getJettonCode(jettonWalletBoc)).EndCell()
	return data
}

func BuildJettonMintPayload(toAddress *address.Address, senderAddress *address.Address,
	amountForward tlb.Coins, jettonAmount tlb.Coins) (_ *cell.Cell, err error) {
	//.storeUint(OPS.InternalTransfer, 32)
	//.storeUint(0, 64)
	//.storeCoins(jettonValue)
	//.storeAddress(null) // TODO FROM?
	//.storeAddress(null) // TODO RESP?
	//.storeCoins(0)
	//.storeBit(false) // forward_payload in this slice, not separate cell
	//.endCell()
	mintMasterMsgPayloadBody := cell.BeginCell().MustStoreUInt(0x178d4519, 32).MustStoreUInt(0, 64).
		MustStoreCoins(jettonAmount.Nano().Uint64()).MustStoreAddr(toAddress).MustStoreAddr(senderAddress).MustStoreCoins(0).
		MustStoreBoolBit(false).EndCell()
	if err != nil {
		return nil, fmt.Errorf("failed to convert JettonMintPayload to cell: %w", err)
	}
	mintPayload := cell.BeginCell().MustStoreUInt(21, 32).MustStoreUInt(0, 64).MustStoreAddr(toAddress).
		MustStoreCoins(amountForward.Nano().Uint64()).MustStoreRef(mintMasterMsgPayloadBody).EndCell()
	return mintPayload, nil
}
