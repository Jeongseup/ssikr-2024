// core/jwt.go
package core

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/multiformats/go-multibase"
)

func VerifyJwt(token string, pbKey *ecdsa.PublicKey) (bool, error) {
	parts := strings.Split(token, ".")
	err := jwt.SigningMethodES256.Verify(strings.Join(parts[0:2], "."), parts[2], pbKey)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Parse VC JWT Claim and Verify VC JWT.
//
// claims의 Issuer에 발급자의 DID가 있다.
// DID를 Resolve해서 DID Document를 받아온다.
// DID도큐먼트의 key ID를 기준으로 public key의 값을 가져와야 하나,
// 여기서는 1개만 존재한다고 가정하고 첫번째를 사용해서 public key를 만들어 사용한다.
func ParseAndVerifyJwtForVC(tokenString string) (bool, *JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {

		//jwt의 암호화 알고리즘이 맞는지 체크한다.
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//발급자의 DID를 추출한다.
		claims := token.Claims.(*JwtClaims)
		issDid := claims.Issuer

		//Resolve한다.
		didDocumentStr, err := ResolveDid(issDid)
		if err != nil {
			log.Printf("Failed to Resolve DID.\nError: %x\n", err)
		}

		//JSON string을 DID Document 객체로 변환한다.
		didDocument, err := NewDIDDocumentForString(didDocumentStr)
		if err != nil {
			log.Printf("Failed generate DID Document from string.\nError: %x\n", err)
		}

		// 첫 번째를 사용한다고 가정한다.
		// TODO: 키 ID(위의 kid)에 해당하는 키 값 구하기. didDocument.findKey(kid)
		kid := token.Header["kid"].(string)
		_ = kid

		pbKeyBaseMultibase := didDocument.VerificationMethod[0].PublicKeyMultibase
		_, bytePubKey, err := multibase.Decode(pbKeyBaseMultibase)
		pbKey, err := x509.ParsePKIXPublicKey(bytePubKey)

		return pbKey, nil
	})

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Vc, claims.Issuer)
		return true, claims, nil
	}

	return false, nil, err
}

func ParseAndVerifyJwtForVP(tokenString string) (bool, *JwtClaimsForVP, error) {
	//개별적으로 내부 VC들을 다시 다 검증해야 한다.
	parseToken, _ := jwt.ParseWithClaims(tokenString, &JwtClaimsForVP{}, func(token *jwt.Token) (interface{}, error) {

		//jwt의 암호화 알고리즘이 맞는지 체크한다.
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			log.Fatalln("unexpected signing method.")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//발급자의 DID를 추출한다.
		claims := token.Claims.(*JwtClaimsForVP)
		issDid := claims.Issuer

		//Resolve한다.
		didDocumentStr, err := ResolveDid(issDid)
		if err != nil {
			log.Printf("Failed to Resolve DID.\nError: %x\n", err)
		}

		//Json string을 DID Document 객체로 생성한다.
		didDocument, err := NewDIDDocumentForString(didDocumentStr)
		if err != nil {
			log.Printf("Failed generate DID Document from string.\nError: %x\n", err)
		}
		// 첫 번째를 사용한다고 가정한다.
		// TODO: 키 ID(위의 kid)에 해당하는 키 값 구하기.
		pbKeyBaseMultibase := didDocument.VerificationMethod[0].PublicKeyMultibase
		_, bytePubKey, err := multibase.Decode(pbKeyBaseMultibase)
		pbKey, err := x509.ParsePKIXPublicKey(bytePubKey)

		return pbKey, nil
	})

	fmt.Println("parseToken: ", parseToken)
	claims, ok := parseToken.Claims.(*JwtClaimsForVP)
	fmt.Println("claims: ", claims)

	if ok && parseToken.Valid {
		if &claims.Vp != nil {
			vpMapClaims := claims.Vp
			vcList := vpMapClaims.VerifiableCredential

			for idx, vcToken := range vcList {
				fmt.Printf("VC[%d]: %s", idx, vcToken)
				verify, _, err := ParseAndVerifyJwtForVC(vcToken)
				if !verify || err != nil {
					log.Printf("Failed to verify VC[%d] in VP.", idx)
					// VC가 한 건이라도 오류라면 바로 리턴해 버린다.
					return false, nil, err
				}
				fmt.Printf("	==> VC[%d] is Verified.\n", idx)
			}
		} else {
			return false, nil, fmt.Errorf("VC is not exist.")
		}

	} else {
		return false, nil, fmt.Errorf("VP is not valid.")
	}

	return true, claims, nil
}
