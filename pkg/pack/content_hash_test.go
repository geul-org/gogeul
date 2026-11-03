//ff:func feature=pack type=helper control=sequence
//ff:what TestContentHash — a fixture document's hash names its h/ copy; odd bytes and non-documents are false
package pack

import (
	"testing"
)

func TestContentHash(t *testing.T) {
	b := fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul")
	h, ok := contentHash(b)
	if !ok || parsePath(hPath("0x00000000", string(h[:]))).space != spaceH {
		t.Fatal(ok)
	}
	if _, ok := contentHash([]byte{1}); ok {
		t.Fatal("odd")
	}
	if _, ok := contentHash([]byte{0, 0}); ok {
		t.Fatal("not a document")
	}
}
