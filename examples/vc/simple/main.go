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
	//verificationMethod := []core.VerificationMethod{
	//	{
	//		Id:                 verificationId,
	//		Type:               "EcdsaSecp256k1VerificationKey2019",
	//		Controller:         issuerDid.String(),
	//		PublicKeyMultibase: issuerKeyEcdsa.PublicKeyMultibase(),
	//	},
	//}
	//didDocument := core.NewDIDDocument(issuerDid.String(), verificationMethod)
	// DID와 DID Document를 VDR에 올려야 하나, 현재 생략.

	// VC 생성.
	// TODO: 나중에 이 파트만 가지고 숙제가 나간다고 함.
	vc, err := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "AlumniCredential"},
		issuerDid.String(),
		map[string]interface{}{
			"id": "1234567890",
			"alumniOf": map[string]interface{}{
				"id": "1234567",
				"name": []map[string]string{
					{
						"value": "Example University",
						"lang":  "en",
					}, {
						"value": "Exemple d'Université",
						"lang":  "fr",
					},
					{
						"value": "예시 대학교",
						"lang":  "kr",
					},
				},
			},
		},
	)

	if err != nil {
		fmt.Println("Failed creation VC.")
		os.Exit(0)
	}

	fmt.Println("my vc:", vc)
	fmt.Println("my verficiationID(from DID):", verificationId)

	// VC에 Issuer의 private key로 서명한다.(JWT 사용)
	token, err := vc.GenerateJWT(verificationId, issuerKeyEcdsa.PrivateKey)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(token)
	fmt.Println("")
	fmt.Println("")

	// 생성된 VC를 검증한다.(public key를 사용해서 검증)
	res, _ := vc.VerifyJwt(token, issuerKeyEcdsa.PublicKey)

	if res {
		fmt.Println("VC is verified.")
	} else {
		fmt.Println("VC is Not verified.")
	}

}
