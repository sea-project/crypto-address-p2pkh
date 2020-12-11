package p2psh

import (
	"encoding/hex"
	ecdsa "github.com/sea-project/crypto-signature-ecdsa"
	"testing"
)

func Test_BTCAddress(t *testing.T) {
	prik, _ := ecdsa.HexToPrvKey("962a3216577de604a0a44086e78960263131b05b92f5ccd3b1d494acf05d3057")
	//k,_ := new(big.Int).SetString("23750643740083697368679458007448773432002777684636444703289980801231536385103",0)
	//prik, pbk := ecdsa.PrivKeyFromBytes(ecc.S256(), k.Bytes())
	//prik, pbk := ecdsa.GenerateKey()
	//pbk,err := ecdsa.HexToPubKey("044230c9577ee5e4d3b947dde274f13978bed3381a170525d8b5a505037ed763a04fae1d7f451aee333260465c524b291dee9854a990e876b7512ab32cc7b13ff6")
	/*if err != nil {
		t.Log(err)
	}*/
	pbk := prik.ToPubKey()
	pubKeyBytes := pbk.SerializeUncompressed()
	t.Log("pubKey", hex.EncodeToString(pubKeyBytes))
	mainAddr := ToBTCAddress(pbk)
	testAddr := ToBTCTestAddress(pbk)
	t.Log(prik.D)
	t.Log(mainAddr)
	t.Log(testAddr)
	addr1 := ToETHAddress(pbk)
	t.Log(addr1.String())
	wif := ecdsa.PrvKeyToWIF(prik, true)
	t.Log(wif)
}
