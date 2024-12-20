// util/utils.go

package util

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcutil/base58"
)

func MakeHash(plain string) []byte {
	digest := sha256.Sum256([]byte(plain))
	return digest[:]
}

func MakeHashBase58(plain string) string {
	return base58.Encode(MakeHash(plain))
}

func MakeHashHex(plain string) string {
	return hex.EncodeToString(MakeHash(plain))
}

func PressKey(msg string) {
	kbReader := bufio.NewReader(os.Stdin)

	fmt.Println(msg)
	kbReader.ReadString('\n')
}

func PrintPrettier(data any) {
	// Convert to pretty JSON
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Failed to generate JSON:", err)
		return
	}

	// Print the pretty JSON
	log.Println(string(prettyJSON))
}

// func MakeHash(plain string) []byte {
// 	digest := sha256.Sum256([]byte(plain))
// 	return digest[:]
// }

// func MakeHashBase58(plain string) string {
// 	return base58.Encode(MakeHash(plain))
// }

// func MakeHashHex(plain string) string {
// 	return hex.EncodeToString(MakeHash(plain))
// }
