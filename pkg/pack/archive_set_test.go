//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveSet — existing sets are returned; new ones join in read order (top first, Issuers in Dir order)
package pack

import "testing"

func TestArchiveSet(t *testing.T) {
	a := &Archive{dir: "wk"}
	b := a.set("0x00000001")
	c := a.set("0x00000002")
	top := a.set("")
	if a.set("0x00000002") != c || len(a.sets) != 3 || a.sets[0] != top || a.sets[1] != b || a.sets[2] != c || b.dir != "wk/0x00000001" {
		t.Fatal(a.sets)
	}
}
