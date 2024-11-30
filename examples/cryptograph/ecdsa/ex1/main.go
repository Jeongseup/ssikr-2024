// example..
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	// go에서는 기본적으로 ECDSA 커브를 제공함.
	// P256 = secp256k1이랑 같음.
	// P256 returns a [Curve] which implements NIST P-256 (FIPS 186-3, section D.2.3), also known as secp256r1 or prime256v1. The CurveParams.Name of this [Curve] is "P-256".

	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Println("err:", err)
	}

	pbKey := &privKey.PublicKey

	fmt.Printf("##### Key Pari #####\n")
	fmt.Printf("Private Key: %x\n", privKey.D)
	// fmt.Printf("Public Key: %x\n", pubKey.X, pubKey.Y)

	fmt.Printf("===== Public Key(X, Y) =====\n")
	fmt.Printf("X=%s Y=%s\n", pbKey.X, pbKey.Y) // 특성상 이렇게 좌표로 구현이 됨.
	// fmt.Printf("  Hex: X=%x Y=%x\n\n", pbKey.X.Bytes(), pbKey.Y.Bytes())
}
