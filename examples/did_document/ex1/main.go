// examples/did_document/ex1/main.go

package main

import (
	"fmt"
	"log"
	"ssikr/core"
)

func main() {
	var method = "ssikr"

	// 1. 키생성(ECDSA)
	kms := new(core.ECDSAManager)
	kms.Generate()

	// 2. DID 생성.
	did, err := core.NewDID(method, kms.PublicKeyMultibase())
	if err != nil {
		log.Printf("Failed to generate DID, error: %v\n", err)
	}

	// 3. DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", did)
	// verification method가 리스트여서 키 자체를 여러개 등록하도록 함.
	verificationMethod := []core.VerificationMethod{
		{
			Id:   verificationId,
			Type: core.VERIFICATION_KEY_TYPE_SECP256K1,
			// DID docs컨트롤하는 주체 여기서는 자기 자신이라서 did 재사용됨
			// authz가 가능한듯?
			Controller:         did.String(),
			PublicKeyMultibase: kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(did.String(), verificationMethod)

	fmt.Println("### Generate DID & DID Document ###")
	fmt.Printf("did => %s\n", did)
	fmt.Printf("did document => %+v\n", didDocument)

}
