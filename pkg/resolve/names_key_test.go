//ff:func feature=resolve type=helper control=sequence
//ff:what TestNamesKey — an Issuer's document list is its reserved document 0xfffffffffffffffe, stored in <I>/docs/
package resolve

import "testing"

func TestNamesKey(t *testing.T) {
	if k := namesKey("0x00000002"); k != (DocKey{Issuer: "0x00000002", Doc: "0xfffffffffffffffe"}) || docFolder(k) != "0x00000002/docs" {
		t.Fatal(k)
	}
}
