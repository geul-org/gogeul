//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteListing — a broken set fails a listing; a remote site has none; a local one lists both sides
package resolve

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSiteListing(t *testing.T) {
	root := packedRoot(t, formsTree(t), false, nil)
	l, err := localSite(root).listing()
	if err != nil {
		t.Fatal(err)
	}
	if o, ok := l.(overlay); !ok || o.loose.root != (Dir{Root: root}).file("") {
		t.Fatal(l)
	}
	s, _ := served(t, root)
	if _, err := s.listing(); err == nil {
		t.Fatal("a remote listing")
	}
	os.WriteFile(filepath.Join(root, ".well-known", "geul", "0x00000000", "latest.txt"), []byte("x"), 0o644)
	if _, err := localSite(root).listing(); err == nil {
		t.Fatal("a broken set")
	}
}
