//ff:func feature=resolve type=helper control=sequence
//ff:what TestKeyDocKey — Issuer and document ID in canonical hex, version ignored
package resolve

import (
	"testing"
)

func TestKeyDocKey(t *testing.T) {
	if d := sampleKey(3, 7).DocKey(); d != (DocKey{Issuer: "0x00000000", Doc: "0xffff000000000003"}) {
		t.Fatalf("got %+v", d)
	}
}
