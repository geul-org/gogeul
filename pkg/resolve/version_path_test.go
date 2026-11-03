//ff:func feature=resolve type=helper control=sequence
//ff:what TestVersionPath — Issuer/document/version.geul
package resolve

import (
	"testing"
)

func TestVersionPath(t *testing.T) {
	if p := versionPath(sampleKey(1, 0).DocKey(), 2); p != "0x00000000/docs/0xFFFF000000000001/0x00000002.geul" {
		t.Fatalf("got %q", p)
	}
}
