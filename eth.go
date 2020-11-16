package p2psh

import (
	sha3 "github.com/sea-project/crypto-hash-sha3"
	ecdsa "github.com/sea-project/crypto-signature-ecdsa"
)

func ToETHAddress(p *ecdsa.PublicKey) Address {
	pubBytes := p.FromECDSAPub()
	i := sha3.Keccak256(pubBytes[1:])[12:]
	return BytesToAddress(i)
}
