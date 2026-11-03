//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestOpenTree — the directory store, or the pack archive with the loose files beside it; a broken pack archive fails
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenTree(t *testing.T) {
	h := newHarness(t)
	h.doc()
	if tr, wk, err := openTree(h.data); err != nil || tr == nil || wk != filepath.Join(h.data, ".well-known", "geul") {
		t.Fatal(err)
	}
	h.packed()
	tr, _, err := openTree(h.data)
	if err != nil {
		t.Fatal(err)
	}
	loose := filepath.Join(h.data, ".well-known", "geul", "0x00000009", "docs", "0x0000000000000001", "versions.geul")
	os.MkdirAll(filepath.Dir(loose), 0o755)
	os.WriteFile(loose, []byte("x"), 0o644)
	if b, err := tr.ReadFile("0x00000009/docs/0x0000000000000001/versions.geul"); err != nil || string(b) != "x" {
		t.Fatal(b, err)
	}
	ents, err := tr.ReadDir(".")
	names := map[string]bool{}
	for _, e := range ents {
		names[e.Name()] = true
	}
	if err != nil || !names["0x00000009"] || names["latest.txt"] {
		t.Fatal(names, err)
	}
	os.WriteFile(filepath.Join(h.data, ".well-known", "geul", "0x00000000", "latest.txt"), []byte("x"), 0o644)
	if _, _, err := openTree(h.data); err == nil {
		t.Fatal("broken")
	}
}
