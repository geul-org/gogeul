//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteOpen — a local site opens a file; a remote one refuses
package resolve

import (
	"io/fs"
	"testing"
)

func TestSiteOpen(t *testing.T) {
	root := packedRoot(t, formsTree(t), false, nil)
	if b, err := fs.ReadFile(localSite(root), issuersPath()); err != nil || len(b) == 0 {
		t.Fatal(err)
	}
	f, err := localSite(root).Open(issuersPath())
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	s, _ := served(t, root)
	if _, err := s.Open(issuersPath()); err == nil {
		t.Fatal("a remote open")
	}
}
