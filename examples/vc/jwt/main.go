package main

import (
	"fmt"
	"os"
	"ssikr/core"
)

// Issuer에 의한 VC 발행 예시.
func main() {
	// 키생성(ECDSA) - 향후 KMS로 대체.
	issuerKeyEcdsa := core.NewEcdsa()

	// DID 생성.
	issuerDid, _ := core.NewDID("ssikr", issuerKeyEcdsa.PublicKeyBase58())

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", issuerDid)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         issuerDid.String(),
			PublicKeyMultibase: issuerKeyEcdsa.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(issuerDid.String(), verificationMethod)
	core.RegisterDid(issuerDid.String(), didDocument.String())

	// VC 생성.
	// 샘플 데이터는 : https://w3c.github.io/vc-data-model/ 참고
	vc, err := core.NewVC(
		"1234567890", // VC에 대한 ID
		[]string{"VerifiableCredential", "AlumniCredential"},
		issuerDid.String(),
		// 크리덴셜 subject인데. 이 subject는 리졸버나 서비스마다 달라질 수 있음
		map[string]interface{}{
			"id": "1234567890", // vc안에 있는 크리덴셜에 대한 id
			"alumniOf": map[string]interface{}{
				"id": "1234567", //
				"name": []map[string]string{
					{
						"value": "Example University",
						"lang":  "en",
					}, {
						"value": "Exemple d'Université",
						"lang":  "fr",
					},
				},
			},
		},
	)

	if err != nil {
		fmt.Println("Failed creation VC.")
		os.Exit(0)
	}

	// VC에 Issuer의 private key로 서명한다.(JWT 사용)
	token, err := vc.GenerateJWT(verificationId, issuerKeyEcdsa.PrivateKey)

	// 생성된 VC를 검증한다.(public key를 사용해서 검증)
	res, _ := vc.VerifyJwt(token, issuerKeyEcdsa.PublicKey)

	if res {
		fmt.Println("VC is verified.")
	} else {
		fmt.Println("VC is Not verified.")
	}

	isVerify, claims, err := core.ParseAndVerifyJwtForVC(token)
	if isVerify {
		fmt.Println("VC is verified.")
		fmt.Printf("Claims: %v\n", claims)
	} else {
		fmt.Println("VC is Not verified.")
	}

}
