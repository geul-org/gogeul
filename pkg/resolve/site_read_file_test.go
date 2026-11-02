//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteReadFile — bytes, fs.ErrNotExist for a missing path, other errors unchanged
package resolve

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSiteReadFile(t *testing.T) {
	tree := formsTree(t)
	root := packedRoot(t, tree, false, nil)
	s := localSite(root)
	want, _ := tree.Fetch(context.Background(), issuersFile)
	if b, err := s.ReadFile(issuersFile); err != nil || string(b) != string(want) {
		t.Fatal(err)
	}
	if _, err := s.ReadFile("0x0000/nothing"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(root, ".well-known", "geul", "0x0000", "latest.txt"), []byte("x"), 0o644)
	if _, err := localSite(root).ReadFile("0x0000/nothing"); !errors.Is(err, pack.ErrCorrupt) {
		t.Fatal(err)
	}
}
