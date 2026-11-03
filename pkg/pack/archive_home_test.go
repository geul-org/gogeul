//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveHome — the top set, else 0x0000's, else the first, else a new top set
package pack

import "testing"

func TestArchiveHome(t *testing.T) {
	a := &Archive{}
	if h := a.home(); h.name != "" || len(a.sets) != 1 {
		t.Fatal(h)
	}
	a = &Archive{sets: []*packSet{{name: "0x00000000"}, {name: "0x00000002"}}}
	if a.home() != a.sets[0] {
		t.Fatal("0x00000000")
	}
	a = &Archive{sets: []*packSet{{name: "0x00000002"}}}
	if a.home() != a.sets[0] {
		t.Fatal("first")
	}
}
