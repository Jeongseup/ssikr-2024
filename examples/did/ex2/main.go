// examples/did/ex2/main.go
package main

import (
	"errors"
	"fmt"
	"log"
	"ssikr/core"
	"ssikr/util"
)

func NewDID(method string, pbKey string) (string, error) {
	if method == "" || pbKey == "" {
		return "", errors.New("parameter is not valid")
	}

	specificIdentifier := util.MakeHashBase58(pbKey)

	// DID:Method:specific
	did := fmt.Sprintf("did:%s:%s", method, specificIdentifier)

	return did, nil
}

func main() {
	var method = "ssikr"

	// key management system shortcut
	kms := new(core.ECDSAManager)
	kms.Generate()

	did, err := NewDID(method, kms.PublicKeyMultibase())

	if err != nil {
		log.Printf("Failed to generate DID, error: %v\n", err)
	}

	fmt.Println("### New DID ###")
	fmt.Printf("did => %s\n", did)
}
