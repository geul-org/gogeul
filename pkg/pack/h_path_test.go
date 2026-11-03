//ff:func feature=pack type=helper control=sequence
//ff:what TestHPath — `<Issuer>/hashes/0x<hex>.geul` and parsePath reads back the Issuer and the 32 hash bytes
package pack

import (
	"strings"
	"testing"
)

func TestHPath(t *testing.T) {
	k := strings.Repeat("\x01", 32)
	p := hPath("0x00000001", k)
	if p != "0x00000001/hashes/0x"+strings.Repeat("01", 32)+".geul" || parsePath(p).hash != k || parsePath(p).issuer != "0x00000001" {
		t.Fatal(p)
	}
}
