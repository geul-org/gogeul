//ff:func feature=resolve type=helper control=sequence
//ff:what TestSubStoreObject — the top set's objects stand at the base, an Issuer set's under its folder
package resolve

import "testing"

func TestSubStoreObject(t *testing.T) {
	if got := (subStore{}).object("latest.txt"); got != "latest.txt" {
		t.Fatal(got)
	}
	if got := (subStore{name: "0x00000002"}).object("m/a.gpm"); got != "0x00000002/m/a.gpm" {
		t.Fatal(got)
	}
}
