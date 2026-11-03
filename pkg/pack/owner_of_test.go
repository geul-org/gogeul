//ff:func feature=pack type=writer control=sequence
//ff:what TestOwnerOf — the Issuer of Issuer paths, "" for top-level and h/ paths
package pack

import "testing"

func TestOwnerOf(t *testing.T) {
	if ownerOf("0x00000000/codecs/0x000003c0/0x00000001.geul") != "0x00000000" || ownerOf("issuers/0x00000001.geul") != "" || ownerOf("h/x.geul") != "" || ownerOf("0x00000000") != "" {
		t.Fatal("ownerOf")
	}
}
