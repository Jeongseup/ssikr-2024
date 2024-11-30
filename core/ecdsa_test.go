// core/ecdsa_test.go

package core

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEcdsaGenerate(t *testing.T) {
	var ecdsa ECDSAManager // ecdsa := new(core.ECDSAManager)
	err := ecdsa.Generate()

	require.NoError(t, err)
}

func TestEcdsaSign(t *testing.T) {
	ecdsa := NewEcdsa()

	msg := "Hello World."
	digest := sha256.Sum256([]byte(msg))
	_, err := ecdsa.Sign(digest[:])
	if err != nil {
		t.Error("Fail to Sign.")
	}
}
