//ff:func feature=pack type=helper control=sequence
//ff:what testArchive — the fixture packed as the Issuer set 0x0000 under a fresh `.well-known/geul/`, opened as an Archive
package pack

import (
	"context"
	"path/filepath"
	"testing"
)

// testArchive returns the opened Archive and its `.well-known/geul/` directory.
func testArchive(t *testing.T) (*Archive, string) {
	t.Helper()
	root := fixture(t)
	wk := filepath.Join(t.TempDir(), ".well-known", "geul")
	if _, err := Build(context.Background(), root, filepath.Join(wk, "0x00000000"), small()); err != nil {
		t.Fatal(err)
	}
	a, err := OpenDir(context.Background(), wk, small())
	if err != nil {
		t.Fatal(err)
	}
	return a, wk
}
