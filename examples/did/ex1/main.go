// examples/did/ex1/main.go
package main

import "fmt"

// DID는 단순 식별자로 W3C에서 제시하는 표준 형식을 따른다.
// https://www.w3.org/TR/did-core/
// ID값은 method가 알아서 관리해야할 몫임.
func main() {
	method := "ssikr"
	specificIdentifier := "abcd1234"

	// [did:DID Method:DID Method-Specific Identifier]
	did := fmt.Sprintf("did:%s:%s", method, specificIdentifier)

	fmt.Printf("DID: %s\n", did)
}
