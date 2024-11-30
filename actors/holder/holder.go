package holder

import (
	"context"
	"fmt"
	"log"
	"os"
	"ssikr/core"
	"ssikr/protos"
	"ssikr/util"
	"time"

	"google.golang.org/grpc"
)

type Holder struct {
	Kms          *core.ECDSAManager
	Did          *core.DID
	DidDocument  *core.DIDDocument
	VCList       []string
	AtomicVCList map[string]string //Atomic VC 저장
}

func (holder *Holder) GenerateDID() {
	// 키생성(ECDSA) - 향후 KMS로 대체.
	holder.Kms = core.NewEcdsa()

	// DID 생성.
	did, _ := core.NewDID("ssikr", holder.Kms.PublicKeyBase58())

	holder.Did = did

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               core.VERIFICATION_KEY_TYPE_SECP256K1,
			Controller:         did.String(),
			PublicKeyMultibase: holder.Kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(did.String(), verificationMethod)
	holder.DidDocument = didDocument
}

type ID struct {
	Name      string
	Mobile    string
	BirthDate string
	Gender    string
}

// ID VC란 주민등록 같은 VC를 의미한다.
// 가장 먼저 발급 받아야하는 것.
func (holder *Holder) GenerateIdentificationVC(id ID) {
	// 원래는 모든 argu에 대해서 validation 필요하겠지만 생략

	// VC 생성.
	vc, _ := core.NewVC(
		"1234567890", // 주민등록 ID 자동으로 생성됨
		[]string{"VerifiableCredential", "SelfCertification"},
		holder.Did.String(), // Issuer did
		map[string]interface{}{
			"id":        "1234567890",
			"name":      id.Name,
			"mobile":    id.Mobile,
			"birthDate": id.BirthDate,
			"gender":    id.Gender,
		},
	)

	fmt.Println("Generated ID VC: ")
	util.PrintPrettier(vc)
	vcJwt, _ := vc.GenerateJWT(holder.DidDocument.VerificationMethod[0].Id, holder.Kms.PrivateKey)
	holder.VCList = append(holder.VCList, vcJwt)
}

func (holder *Holder) GenerateFirstVC() {
	// VC 생성.
	vc, _ := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "SelfCertification"},
		holder.Did.String(), // Issuer did
		map[string]interface{}{
			"id":        "1234567890",
			"name":      "SON JEONGSEUP",
			"mobile":    "010-1234-1234",
			"birthDate": "2000-01-01",
			"gender":    "M",
		},
	)

	vcJwt, _ := vc.GenerateJWT(holder.DidDocument.VerificationMethod[0].Id, holder.Kms.PrivateKey)
	holder.VCList = append(holder.VCList, vcJwt)
}

func (holder *Holder) GenerateVPForID() (string, error) {
	VCForID := holder.VCList[0]

	vp, err := core.NewVP(
		"12345678901111",
		[]string{"VerifiablePresentaion"},
		holder.Did.String(),
		[]string{VCForID},
	)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	vpToken := vp.GenerateJWT(holder.DidDocument.VerificationMethod[0].Id, holder.Kms.PrivateKey)

	return vpToken, nil
}

func (holder *Holder) GenerateVP() (string, error) {
	vcList := holder.VCList

	vp, err := core.NewVP(
		"12345678901111",
		[]string{"VerifiablePresentaion"},
		holder.Did.String(),
		vcList,
	)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	vpToken := vp.GenerateJWT(holder.DidDocument.VerificationMethod[0].Id, holder.Kms.PrivateKey)

	return vpToken, nil
}

// 실제로 방문하거나 전화를 통해서만 발급가능한 VC
func (holder *Holder) RequestVCToRootOfTrustIssuer(name, mobile, birthDate, gender string) error {
	conn, err := grpc.Dial("localhost:1120", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("RootOfTrustIssuer not connect: %v", err)
		return err
	}
	defer conn.Close()

	c := protos.NewSimpleIssuerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := c.IssueIDVC(ctx, &protos.MsgRequestIDVC{
		Did:       holder.Did.String(),
		Name:      name,
		Mobile:    mobile,
		BirthDate: birthDate,
		Gender:    gender,
	})
	if err != nil {
		log.Printf("could not request: %v", err)
		return err
	}

	fmt.Printf("RootOfTrustIssuer's response: %s\n", res.Result)
	fmt.Printf("RootOfTrustIssuer's response VC: %s\n", res.Vc)
	if res.Result == "OK" {
		holder.VCList = append(holder.VCList, res.Vc)
	}

	return nil
}

func (holder *Holder) RequestVCToUniversityIssuer(vpToken string) error {
	conn, err := grpc.Dial("localhost:1121", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("UniversityIssuer not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := protos.NewSimpleIssuerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.IssueSimpleVC(ctx, &protos.MsgRequestVC{
		Did: holder.Did.String(),
		Vp:  vpToken,
	})
	if err != nil {
		log.Printf("could not request: %v", err)
		return err
	}

	fmt.Printf("UniversityIssuer's response: %s\n", res.Result)
	fmt.Printf("UniversityIssuer's response VC: %s\n", res.Vc)
	if res.Result == "OK" {
		holder.VCList = append(holder.VCList, res.Vc)
	}

	return nil
}

func (holder *Holder) RequestVCToCompanyIssuer(vpToken string) error {
	conn, err := grpc.Dial("localhost:1122", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("CompanyIssuer not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := protos.NewSimpleIssuerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.IssueSimpleVC(ctx, &protos.MsgRequestVC{
		Did: holder.Did.String(),
		Vp:  vpToken,
	})
	if err != nil {
		log.Printf("could not request: %v", err)
		return err
	}

	fmt.Printf("CompanyIssuer's response: %s\n", res.Result)
	fmt.Printf("CompanyIssuer's response VC: %s\n", res.Vc)
	if res.Result == "OK" {
		holder.VCList = append(holder.VCList, res.Vc)
	}

	return nil
}

func (holder *Holder) RequestVCToBankIssuer(vpToken string) error {
	conn, err := grpc.Dial("localhost:1123", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("BankIssuer not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := protos.NewMultipleIssuerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.IssueMultipleVC(ctx, &protos.MsgRequestMultipleVC{
		Did: holder.Did.String(),
		Vp:  vpToken,
	})
	if err != nil {
		log.Printf("could not request: %v", err)
		return err
	}

	fmt.Printf("BankIssuer's response: %s\n", res.Result)

	if res.Result == "OK" {
		for _, vc := range res.Vc {
			fmt.Printf("BankIssuer's response VC: %s\n", vc)
			holder.VCList = append(holder.VCList, vc)
		}
	}

	return nil
}

func (holder *Holder) RequestVCToAtomicUniversityIssuer(vpToken string) error {
	conn, err := grpc.Dial("localhost:1124", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("AtomicUniversityIssuer not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := protos.NewAtomicIssuerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.IssueAtomicVC(ctx, &protos.MsgRequestAtomicVC{
		Did: holder.Did.String(),
		Vp:  vpToken,
	})
	if err != nil {
		log.Printf("could not request: %v", err)
		return err
	}

	fmt.Printf("AtomicUniversity's response: %s\n", res.Result)

	if res.Result == "OK" {
		for _, vc := range res.Vcs {
			fmt.Printf("AtomicUniversity's response VC Name: %s\n", vc.Name)
			fmt.Printf("AtomicUniversity's response VC Token: %s\n", vc.Token)
			holder.AtomicVCList[vc.Name] = vc.Token
		}
	}

	return nil
}

func (holder *Holder) PrintAtomicVC() {
	if holder.AtomicVCList == nil {
		return
	}
	fmt.Println("<< Atomic VC List >>")
	idx := 1
	for key, _ := range holder.AtomicVCList {
		fmt.Println(idx, ". ", key)
		idx++
	}
}

func (holder *Holder) PrintVCTokens() {
	if holder.VCList == nil {
		return
	}
	// fmt.Println("<< VC List >>")
	for idx, vc := range holder.VCList {
		fmt.Printf("[%d]VC: %+v\n", idx, vc)
	}
}

func (holder *Holder) PrintVCDetails() {
	if holder.VCList == nil {
		return
	}
	// fmt.Println("<< VC Detail List >>")
	for idx, vcToken := range holder.VCList {
		isVerify, claims, err := core.ParseAndVerifyJwtForVC(vcToken)
		if (!isVerify) || (err != nil) {
			fmt.Println("VC is Not verified. or got err..", err)
		}
		fmt.Printf("[%d]VC: \n", idx)
		util.PrintPrettier(claims)
	}
}
