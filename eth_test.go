package p2psh

import (
	ecdsa "github.com/sea-project/crypto-signature-ecdsa"
	stdlib "github.com/sea-project/stdlib-hex-conversion"
	"math/big"
	"testing"
)

func Test_Address(t *testing.T) {
	prik, pbk := ecdsa.GenerateKey()

	addr := ToETHAddress(pbk).ToICAP("sea", "0000")
	t.Log(prik.D)
	t.Log(addr)
}

func Test_Change(t *testing.T) {
	key, _ := big.NewInt(0).SetString("49016675427890240271764302863626314979539702441087103095488164811326728265118", 10)
	to36 := stdlib.DECToBHex(key)
	t.Log(to36)
	a1 := big.NewInt(0).SetBytes([]byte(to36))
	t.Log(a1)
}
