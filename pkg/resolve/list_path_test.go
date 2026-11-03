//ff:func feature=resolve type=helper control=sequence
//ff:what TestListPath — Issuer/document/versions.geul
package resolve

import (
	"testing"
)

func TestListPath(t *testing.T) {
	if p := listPath(sampleKey(1, 0).DocKey()); p != "0x00000000/docs/0xFFFF000000000001/versions.geul" {
		t.Fatalf("got %q", p)
	}
}
