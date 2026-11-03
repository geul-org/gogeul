//ff:func feature=server type=handler control=sequence
//ff:what TestTreeSpec — no specs gives an empty Spec; a broken spec document gives none
package server

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTreeSpec(t *testing.T) {
	root := treeFixture(t)
	if s := treeSpec(treeOf(root)); s == nil || len(s) != 0 {
		t.Fatalf("empty: %v", s)
	}
	dir := filepath.Join(root, ".well-known", "geul", "0x00000000", "codecs", "0x00000b80")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "versions.geul"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if s := treeSpec(treeOf(root)); s != nil {
		t.Fatalf("broken: %v", s)
	}
}
