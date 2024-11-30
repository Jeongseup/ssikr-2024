package main

import (
	"fmt"
	"ssikr/actors/holder"
	"ssikr/core"
	"ssikr/util"
)

func main() {
	fmt.Println("### Start HOLDER's Wallet ###")
	// New Holder
	hldr := new(holder.Holder)
	hldr.AtomicVCList = make(map[string]string)

	fmt.Println("1. Holder의 DID를 생성")
	hldr.GenerateDID()
	fmt.Printf("Holder's DID: %s\n", hldr.Did.String())
	fmt.Println("Holder's DID Document:")
	util.PrintPrettier(hldr.DidDocument)

	fmt.Println("2. DID를 VDR에 등록합니다.")
	core.RegisterDid(hldr.Did.String(), hldr.DidDocument.String())

	// 먼저 RootOfTrust VC 발급
	fmt.Println("3. 최초 RootOfTrust VC(ID)를 발급합니다.")
	myName, myMobile, myBirthDate, myGender := "SON JEONGSEUP", "010 1234 1234", "2024-11-30", "M"
	// RootOfTrust에게 ID VC 인증을 요청
	hldr.RequestVCToRootOfTrustIssuer(myName, myMobile, myBirthDate, myGender)

	// UniversityIssuer에게 졸업증명 VC를 요청한다.
	// util.PressKey("4. UniversityIssuer에게 졸업증명 VC를 요청한다. [아무키나 입력하세요.]")
	fmt.Println("4. UniversityIssuer에게 졸업증명 VC를 요청")
	vpToken, _ := hldr.GenerateVP()
	fmt.Printf("VP Token: %s\n", vpToken)
	hldr.RequestVCToUniversityIssuer(vpToken)

	hldr.PrintVCTokens()
	hldr.PrintVCDetails()
}
