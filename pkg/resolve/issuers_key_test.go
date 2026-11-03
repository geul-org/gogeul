//ff:func feature=resolve type=helper control=sequence
//ff:what TestIssuersKey — the Issuer list is geul.org 0x00000000's reserved document 0xffffffffffffffff, stored in issuers/
package resolve

import "testing"

func TestIssuersKey(t *testing.T) {
	if k := issuersKey(); k != (DocKey{Issuer: "0x00000000", Doc: "0xffffffffffffffff"}) || docFolder(k) != "issuers" {
		t.Fatal(k)
	}
}
