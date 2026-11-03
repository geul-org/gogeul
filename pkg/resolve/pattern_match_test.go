//ff:func feature=resolve type=helper control=sequence
//ff:what TestPatternMatch — whole Issuer, one document, `*`; other Issuers and documents do not match
package resolve

import (
	"testing"
)

func TestPatternMatch(t *testing.T) {
	d := DocKey{Issuer: "0x00000000", Doc: "0xffff000000000001"}
	if !(Pattern{Issuer: "0x00000000"}).Match(d) || !(Pattern{Issuer: "*"}).Match(d) || !(Pattern{Issuer: "0x00000000", Doc: d.Doc}).Match(d) {
		t.Fatal("want match")
	}
	if (Pattern{Issuer: "0x00000002"}).Match(d) || (Pattern{Issuer: "0x00000000", Doc: "0xffff000000000002"}).Match(d) {
		t.Fatal("want no match")
	}
}
