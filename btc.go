package p2psh

import (
	"crypto/sha256"
	base58 "github.com/sea-project/crypto-codec-base58"
	ecdsa "github.com/sea-project/crypto-signature-ecdsa"
	"golang.org/x/crypto/ripemd160"
)

const (
	mainnetVersion     = byte(0x00) // 定义版本号，一个字节[mainnet]
	addressChecksumLen = 4          // 定义checksum长度为四个字节
)

func ToBTCAddress(p *ecdsa.PublicKey) string {
	pubKeyBytes := p.SerializeCompressed()
	ripemd160Hash := Ripemd160Hash(pubKeyBytes)
	versionRipemd160hash := append([]byte{mainnetVersion}, ripemd160Hash...)
	//调用CheckSum方法返回前四个字节的checksum
	checkSumBytes := CheckSum(versionRipemd160hash)
	//将version+Pub Key hash+ checksum生成25个字节
	bytes := append(versionRipemd160hash, checkSumBytes...)
	return base58.Encode(bytes, base58.BitcoinAlphabet)
}

// 对公钥进行sha256散列和ripemd160散列,获得publickeyHash
func Ripemd160Hash(publicKey []byte) []byte {
	//将传入的公钥进行256运算，返回256位hash值
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hash := hash256.Sum(nil)
	//将上面的256位hash值进行160运算，返回160位的hash值
	ripemd160 := ripemd160.New()
	ripemd160.Write(hash)

	return ripemd160.Sum(nil) //返回Pub Key hash
}

// sha256(sha256(versionPublickeyHash))  取最后4个字节的值
func CheckSum(payload []byte) []byte {
	//这里传入的payload其实是version+Pub Key hash，对其进行两次256运算
	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:addressChecksumLen] //返回前四个字节，为CheckSum值
}
