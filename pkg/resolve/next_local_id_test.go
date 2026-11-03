//ff:func feature=resolve type=resolver control=sequence
//ff:what TestNextLocalID — one above the document's largest LOCAL, kept at the earlier value when that is higher, 1 for none
package resolve

import "testing"

func TestNextLocalID(t *testing.T) {
	b := codecDocOf(0x0A04, 4, 4)
	if n := nextLocalID(0, b); n < 2 {
		t.Fatalf("doc: %d", n)
	}
	if n := nextLocalID(1000, b); n != 1000 {
		t.Fatalf("kept: %d", n)
	}
	if n := nextLocalID(0, nil); n != 1 {
		t.Fatalf("none: %d", n)
	}
}
