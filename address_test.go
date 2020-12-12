package p2psh

import "testing"

func TestAddress_ToICAP(t *testing.T) {
	address := "0xE0Cc8F88e6Ce7e384398996ecA7978540534C18B"
	addr := HexToAddress(address).ToICAP("sea", "0000")
	t.Log(addr)
}