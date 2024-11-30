package rootoftruestissuer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"ssikr/core"
	"ssikr/protos"
	"ssikr/util"
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

func (server *Server) IssueIDVC(_ context.Context, msg *protos.MsgRequestIDVC) (*protos.MsgResponseVC, error) {
	log.Println("called IssueIDVC method")
	log.Println("ROT VC 발급중!!!!")
	log.Printf("대면 인증: %s씨 %s은 당신의 DID가 맞습니까?", msg.GetDid(), msg.String())
	util.PressKey("ROT는 대면인증방법을 통해 체크하고 발급해준다")

	vcToken, err := server.Issuer.GenerateIDVC(msg)
	if err != nil {
		response := new(protos.MsgResponseVC)
		response.Result = "FAIL"
		return response, nil
	}

	response := new(protos.MsgResponseVC)
	response.Result = "OK"
	response.Vc = vcToken
	return response, nil

}

func (issuer *Issuer) GenerateIDVC(msg *protos.MsgRequestIDVC) (string, error) {
	// VC 생성.
	vc, err := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "ResidentRegistrationCredential"},
		issuer.did.String(),
		map[string]interface{}{
			"did":       "1234567890",
			"name":      msg.Name,
			"mobile":    msg.Mobile,
			"birthDate": msg.BirthDate,
			"gender":    msg.Gender,
		},
	)

	if err != nil {
		return "", errors.New("Failed creation VC.")
	}

	log.Println("Generated ID VC: ")
	util.PrintPrettier(vc)

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
