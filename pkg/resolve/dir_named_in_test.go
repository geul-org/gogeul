//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirNamedIn — ID folders of the given width, without names; files, other widths and a missing folder are left out
package resolve

import "testing"

func TestDirNamedIn(t *testing.T) {
	f := newTypedFixture(t)
	got := f.d.namedIn("0x00000000/docs", 4)
	if len(got) != 1 || got[0].ID != "0xffff000000000007" || got[0].Names != nil {
		t.Fatal(got)
	}
	if got := f.d.namedIn("0x00000000/codecs", 2); len(got) != 1 || got[0].ID != "0x00000980" {
		t.Fatal(got)
	}
	if got := f.d.namedIn("0x00000000/docs", 2); len(got) != 0 {
		t.Fatal(got)
	}
	if got := f.d.namedIn("0x00000009/docs", 4); got != nil {
		t.Fatal(got)
	}
}
