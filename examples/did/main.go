// examples/did/main.go
package main

import (
	"fmt"
	core "ssikr/core"
)

func main() {
	did, _ := core.NewDID("test", "12345")

	fmt.Printf("DID: [%s]", did.String())

}
