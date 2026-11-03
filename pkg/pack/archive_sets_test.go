//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveSets — sets with a generation, in read order; a new set without one is left out
package pack

import "testing"

func TestArchiveSets(t *testing.T) {
	a, _ := testArchive(t)
	a.set("0x00000002")
	s := a.Sets()
	if len(s) != 1 || s[0].Name != "0x00000000" || s[0].Generation != 1 || s[0].Objects != 2 || len(s[0].Manifest) != 64 {
		t.Fatal(s)
	}
}
