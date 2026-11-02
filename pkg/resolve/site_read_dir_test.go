//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestSiteReadDir — a local site lists packed and loose Issuers; a remote one refuses
package resolve

import (
	"slices"
	"testing"
)

func TestSiteReadDir(t *testing.T) {
	root := packedRoot(t, formsTree(t), false, nil)
	if err := (Dir{Root: root}).write("0x0009/0x0000000000000001/versions.geul", []byte("x")); err != nil {
		t.Fatal(err)
	}
	ents, err := localSite(root).ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, e := range ents {
		names = append(names, e.Name())
	}
	if !slices.Equal(names, []string{"0x0000", "0x0001", "0x0009", "h", "issuers.geul"}) {
		t.Fatal(names)
	}
	s, _ := served(t, root)
	if _, err := s.ReadDir("."); err == nil {
		t.Fatal("a remote listing")
	}
}
