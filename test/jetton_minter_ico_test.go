package jetton

import (
	"fmt"
	"github.com/e9571/lib1"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton/jetton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"log"
	"testing"
	"ton-demo/utils"
)

func TestDeployJettonMintIco(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())
	msgBody := cell.BeginCell().EndCell()
	fmt.Println("Deploying jetton ico contract to testnet...")
	jettonWalletBoc := "b5ee9c7241021201000334000114ff00f4a413f4bcf2c80b0102016202110202cc03060201d4040500c30831c02497c138007434c0c05c6c2544d7c0fc02f83e903e900c7e800c5c75c87e800c7e800c1cea6d0000b4c7e08403e29fa954882ea54c4d167c0238208405e3514654882ea58c511100fc02780d60841657c1ef2ea4d67c02b817c12103fcbc2000113e910c1c2ebcb85360020148070e020120080a01f100f4cffe803e90087c007b51343e803e903e90350c144da8548ab1c17cb8b04a30bffcb8b0950d109c150804d50500f214013e809633c58073c5b33248b232c044bd003d0032c032483e401c1d3232c0b281f2fff274013e903d010c7e800835d270803cb8b11de0063232c1540233c59c3e8085f2dac4f3200900ae8210178d4519c8cb1f19cb3f5007fa0222cf165006cf1625fa025003cf16c95005cc2391729171e25008a813a08208e4e1c0aa008208989680a0a014bcf2e2c504c98040fb001023c85004fa0258cf1601cf16ccc9ed5403f73b51343e803e903e90350c0234cffe80145468017e903e9014d6f1c1551cdb5c150804d50500f214013e809633c58073c5b33248b232c044bd003d0032c0327e401c1d3232c0b281f2fff274140371c1472c7cb8b0c2be80146a2860822625a020822625a004ad8228608239387028062849f8c3c975c2c070c008e00b0c0d00705279a018a182107362d09cc8cb1f5230cb3f58fa025007cf165007cf16c9718010c8cb0524cf165006fa0215cb6a14ccc971fb0010241023000e10491038375f040076c200b08e218210d53276db708010c8cb055008cf165004fa0216cb6a12cb1f12cb3fc972fb0093356c21e203c85004fa0258cf1601cf16ccc9ed540201200f1000db3b51343e803e903e90350c01f4cffe803e900c145468549271c17cb8b049f0bffcb8b0a0823938702a8005a805af3cb8b0e0841ef765f7b232c7c572cfd400fe8088b3c58073c5b25c60063232c14933c59c3e80b2dab33260103ec01004f214013e809633c58073c5b3327b55200083200835c87b51343e803e903e90350c0134c7e08405e3514654882ea0841ef765f784ee84ac7cb8b174cfcc7e800c04e81408f214013e809633c58073c5b3327b5520001ba0f605da89a1f401f481f481a861f0a7d84b"
	jettonMintIcoBoc := "b5ee9c7241020b010001f5000114ff00f4a413f4bcf2c80b0102016202080202cc030703f7d80e8698180b8d8492f81f07d201876a2687d007d206a6a1812e38047221ac1044c4b4028b350906100797026381041080bc6a28ce4658fe59f917d017c14678b13678b10fd0165806493081b2044780402502189e428027d012c678b666664f6aa701b02698fe99fc00aa9185d718141083deecbef09dd71812f83c04050600606c215131c705f2e04902fa40fa00d43020d08060d721fa00302710345042f008a05023c85004fa0258cf16ccccc9ed5400fc01fa00fa40f82854120970542013541403c85004fa0258cf1601cf16ccc922c8cb0112f400f400cb00c9f9007074c8cb02ca07cbffc9d05006c705f2e04a13a1034145c85004fa0258cf16ccccc9ed5401fa403020d70b01c3008e1f8210d53276db708010c8cb055003cf1622fa0212cb6acb1fcb3fc98042fb00915be20008840ff2f00093dfc142201b82a1009aa0a01e428027d012c678b00e78b666491646580897a007a00658064907c80383a6465816503e5ffe4e83bc00c646582ac678b28027d0109e5b589666664b8fd80402037a60090a007dadbcf6a2687d007d206a6a183618fc1400b82a1009aa0a01e428027d012c678b00e78b666491646580897a007a00658064fc80383a6465816503e5ffe4e840001faf16f6a2687d007d206a6a183faa9040521c9d9d"
	addr, _, _, err := w.DeployContractWaitTransaction(ctx, tlb.MustFromTON("0.1"),
		msgBody, getJettonCode(jettonMintIcoBoc), getJettonContractDataJson(w.WalletAddress(), jettonWalletBoc))
	if err != nil {
		panic(err)
	}
	//EQCU9EivCHA7bTIPUoQMWDSwAvyqaqWbZ9B351DysLFOA7Mb
	fmt.Println("Deployed contract addr:", addr.String())
}

func TestMintIcoJettonWithTon(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())
	mintAddr := address.MustParseAddr("EQCU9EivCHA7bTIPUoQMWDSwAvyqaqWbZ9B351DysLFOA7Mb")
	fmt.Println("Minting jettons...")
	mint := wallet.SimpleMessageAutoBounce(mintAddr, tlb.MustFromTON("0.33"), nil)
	_, _, err := w.SendWaitTransaction(ctx, mint)
	if err != nil {
		panic(err)
	}
}

func TestMintIcoJetton(t *testing.T) {
	api, ctx := utils.GetTestApiAndClient(true)
	w := utils.GetWallet(api)
	log.Println("Deploy wallet:", w.WalletAddress().String())
	mintAddr := address.MustParseAddr("EQCU9EivCHA7bTIPUoQMWDSwAvyqaqWbZ9B351DysLFOA7Mb")
	token := jetton.NewJettonMasterClient(api, mintAddr)

	jettonData, err := token.GetJettonData(ctx)
	if err != nil {
		panic(err)
	}
	utils.LogOut("info,", "jettonData:"+lib1.Json_Package(jettonData))
	toAddress := address.MustParseAddr("UQDUkb17UDGmYyQXst7oiPWdALTRlECaxgmlPioeGi7nvlPT")
	mintData, err := BuildJettonMintIcoPayload(toAddress, w.WalletAddress(), tlb.MustFromTON("0.015"), tlb.MustFromTON("0.009"))
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

func BuildJettonMintIcoPayload(toAddress *address.Address, senderAddress *address.Address,
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
