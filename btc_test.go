package p2psh

import (
	"encoding/hex"
	ecdsa "github.com/sea-project/crypto-signature-ecdsa"
	"testing"
)

func Test_BTCAddress(t *testing.T) {
	prik, _ := ecdsa.HexToPrvKey("28ea039252a3c0b5f3ec2d92f664011561ccf69f434512f20d0daa5fb2a34931")
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
	addr := ToBTCAddress(pbk)
	t.Log(prik.D)
	t.Log(addr)
	addr1 := ToETHAddress(pbk)
	t.Log(addr1.String())
}
