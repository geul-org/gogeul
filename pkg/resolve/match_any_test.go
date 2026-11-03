//ff:func feature=resolve type=helper control=sequence
//ff:what TestMatchAny — true when one pattern matches, false for none or an empty list
package resolve

import (
	"testing"
)

func TestMatchAny(t *testing.T) {
	d := DocKey{Issuer: "0x00000000", Doc: "0xFFFF000000000001"}
	if !matchAny([]Pattern{{Issuer: "0x00000002"}, {Issuer: "0x00000000"}}, d) {
		t.Fatal("want match")
	}
	if matchAny([]Pattern{{Issuer: "0x00000002"}}, d) || matchAny(nil, d) {
		t.Fatal("want no match")
	}
}
