package core

import (
	"crypto/ecdsa"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	// "github.com/mohae/deepcopy"
	"github.com/getlantern/deepcopy"
)

// https://www.w3.org/TR/vc-data-model
// https://www.w3.org/TR/vc-data-model/#json-web-token
type VC struct {
	// Mendatory
	Context []string `json:"@context"`

	Id                string                 `json:"id,omitempty"`
	Type              []string               `json:"type,omitempty"`
	Issuer            string                 `json:"issuer,omitempty"`
	IssuanceDate      string                 `json:"issuanceDate,omitempty"`
	CredentialSubject map[string]interface{} `json:"credentialSubject,omitempty"`
	Proof             *Proof                 `json:"proof,omitempty"`
}

type Proof struct {
	Type               string `json:"type,omitempty"`
	Created            string `json:"created,omitempty"`
	ProofPurpose       string `json:"proofPurpose,omitempty"`
	VerificationMethod string `json:"verificationMethod,omitempty"`
	ProofValue         string `json:"proofValue,omitempty"`
	Jws                string `json:"jws,omitempty"`
}

// JWT를 위한 claim
type JwtClaims struct {
	jwt.StandardClaims

	Nonce string
	Vc    VC `json:"vc,omitempty"`
}

func NewVC(id string, typ []string, issuer string, credentialSubject map[string]interface{}) (*VC, error) {
	newVC := &VC{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/v2",
		},
		Id:                id,
		Type:              typ,
		Issuer:            issuer,
		IssuanceDate:      time.Now().Format(time.RFC3339), //"2010-01-01T19:23:24Z",
		CredentialSubject: credentialSubject,
	}
	return newVC, nil
}

type VCInterface interface {
	GenerateJWT() string
	VerifyJwt() (bool, error)
}

// VC를 JTW로 생성하고 string으로 반환한다.
// JTW의 경우 JWS로 증명되기에 Proofs를 빼고, JWT와 중복되는 properties를 제거한다.
func (vc *VC) GenerateJWT(verificationId string, pvKey *ecdsa.PrivateKey) (string, error) {
	aud := ""
	exp := time.Now().Add(time.Minute * 5).Unix()       //만료 시간. 현재 + 5분
	jti := "1112342"                                    // JWT ID
	t, err := time.Parse(time.RFC3339, vc.IssuanceDate) //unixtime으로 바꾸기 위해.
	iat := t.Unix()
	nbf := iat
	iss := vc.Issuer
	sub := "Verifiable Credential"

	// Proof를 제거하고 JWT를 만들기 위해 복제한다.
	vcTmp := new(VC)
	deepcopy.Copy(vcTmp, vc)
	vcTmp.Proof = nil

	jwtClaims := JwtClaims{
		jwt.StandardClaims{
			Audience:  aud,
			ExpiresAt: exp,
			Id:        jti,
			IssuedAt:  iat,
			Issuer:    iss,
			NotBefore: nbf,
			Subject:   sub,
		},
		"qwasd!234",
		*vcTmp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwtClaims)
	token.Header["kid"] = verificationId

	tokenString, err := token.SignedString(pvKey)

	if err != nil {

	}

	return tokenString, nil
}

func (vc *VC) VerifyJwt(token string, pbKey *ecdsa.PublicKey) (bool, error) {
	parts := strings.Split(token, ".")
	err := jwt.SigningMethodES256.Verify(strings.Join(parts[0:2], "."), parts[2], pbKey)
	if err != nil {
		return false, err
	}

	return true, nil
}
