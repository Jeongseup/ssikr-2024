package UniversityIssuer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"ssikr/core"
	"ssikr/protos"
)

type Server struct {
	protos.UnimplementedSimpleIssuerServer

	Issuer *Issuer
}

type Issuer struct {
	kms         *core.ECDSAManager
	did         *core.DID
	didDocument *core.DIDDocument

	CredentialSubjectJsonFilePath string
}

func (issuer *Issuer) GenerateDID() {
	// 키생성(ECDSA)
	issuer.kms = core.NewEcdsa()

	// DID 생성.
	issuerDid, _ := core.NewDID("ssikr", issuer.kms.PublicKeyBase58())

	issuer.did = issuerDid

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", issuerDid)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               core.VERIFICATION_KEY_TYPE_SECP256K1,
			Controller:         issuerDid.String(),
			PublicKeyMultibase: issuer.kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(issuerDid.String(), verificationMethod)
	issuer.didDocument = didDocument

	registerDid(issuerDid.String(), didDocument)
}

func (server *Server) IssueSimpleVC(_ context.Context, msg *protos.MsgRequestVC) (*protos.MsgResponseVC, error) {
	if msg.Vp == "" || msg.Vp == "NONE" {
		response := new(protos.MsgResponseVC)
		response.Result = "FAIL"
		response.Vc = "VP is invalid"
		return response, nil
	}

	isVerify, claims, err := core.ParseAndVerifyJwtForVP(msg.Vp)
	if !isVerify || err != nil {
		fmt.Println("VP is NOT verified.")
		return nil, errors.New(fmt.Sprintf("VP is invalid: %s", err))
	}

	for i, vc := range claims.Vp.VerifiableCredential {
		isVerify, claims, err := core.ParseAndVerifyJwtForVC(vc)
		if !isVerify || err != nil {
			fmt.Println("VC #", i, " is NOT verified.")
			return nil, errors.New(fmt.Sprintf("VC is invalid: %s", err))
		}
		fmt.Println("### VC is verified and keep issuing Universtiy VC... ###")
		vcClaims := claims.Vc.CredentialSubject
		// fmt.Printf("VC claims:: %+v", vcClaims)

		// Extract the value associated with the "name" key
		name, exists := vcClaims["name"]
		if !exists {
			fmt.Println("Key 'name' not found in the map")
			return nil, errors.New(fmt.Sprintf("VP is invalid: %s", err))
		}

		vcToken, err := server.Issuer.GenerateSampleVC(name.(string))
		if err != nil {
			fmt.Println("Failed to generate University VC")
			return nil, errors.New(fmt.Sprintf("VP is invalid: %s", err))
		}

		response := new(protos.MsgResponseVC)
		response.Result = "OK"
		response.Vc = vcToken
		return response, nil
	}

	return nil, errors.New("Error")
}

func (issuer *Issuer) GenerateSampleVC(name string) (string, error) {
	var credentialSubject map[string]interface{}
	vcData := make(map[string]interface{})
	vcData["name"] = name
	credentialSubject = vcData

	// VC 생성.
	vc, err := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "IncheonNationalUniversity"},
		issuer.did.String(),
		credentialSubject,
	)

	if err != nil {
		return "", errors.New("Failed creation VC.")
	}

	// VC에 Issuer의 private key로 서명한다.(JWT 사용)
	token, err := vc.GenerateJWT(issuer.didDocument.VerificationMethod[0].Id, issuer.kms.PrivateKey)

	return token, nil
}

func registerDid(did string, document *core.DIDDocument) error {
	err := core.RegisterDid(did, document.String())
	if err != nil {
		return err
	}

	return nil
}

func loadJson(path string) map[string]interface{} {
	jsonData := make(map[string]interface{})

	data, err := os.Open(path)
	if err != nil {
		return nil
	}

	byteValue, _ := ioutil.ReadAll(data)

	json.Unmarshal(byteValue, &jsonData)

	return jsonData
}
