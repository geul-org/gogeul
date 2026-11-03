//ff:func feature=build type=command control=iteration dimension=1
//ff:what TestTree — bytes, pages, listings and scripts at the public paths; output inside the source, a bad document and an unwritable output fail
package build

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTree(t *testing.T) {
	root := fixture(t)
	dst := t.TempDir()
	c, err := Tree(root, dst)
	if err != nil || c != (Count{Docs: 2, Dirs: 4}) {
		t.Fatalf("count: %+v %v", c, err)
	}
	out := filepath.Join(dst, ".well-known", "geul")
	src, _ := os.ReadFile(filepath.Join(root, ".well-known", "geul", "h", "a.geul"))
	if b, _ := os.ReadFile(filepath.Join(out, "h", "a.geul")); !bytes.Equal(b, src) {
		t.Fatal("bytes not copied")
	}
	for p, want := range map[string]string{
		"h/a":                      "browseBtn",
		"0x00000000/x/0x00000001":      "browseBtn",
		"index.html":               `href="/.well-known/geul/h/"`,
		"h/index.html":             `href="/.well-known/geul/h/a"`,
		"h/index.json":             `"URL":"/.well-known/geul/h/a"`,
		"0x00000000/x/index.json":      `"URL":"/.well-known/geul/0x00000000/x/0x00000001"`,
		"_static/cytoscape.min.js": "cytoscape",
	} {
		if b, err := os.ReadFile(filepath.Join(out, filepath.FromSlash(p))); err != nil || !strings.Contains(string(b), want) {
			t.Fatalf("%s: %v", p, err)
		}
	}
	if _, err := os.Stat(filepath.Join(out, "h", "notes.txt")); err == nil {
		t.Fatal("non-.geul file copied")
	}
	if _, err := Tree(root, filepath.Join(root, ".well-known", "geul", "h")); err == nil {
		t.Fatal("inside source: want error")
	}
	os.WriteFile(filepath.Join(root, ".well-known", "geul", "h", "bad.geul"), []byte{1, 2}, 0o644)
	if _, err := Tree(root, t.TempDir()); err == nil || !strings.Contains(err.Error(), "h/bad.geul") {
		t.Fatalf("bad document: %v", err)
	}
	os.Remove(filepath.Join(root, ".well-known", "geul", "h", "bad.geul"))
	blocked := t.TempDir()
	os.MkdirAll(filepath.Join(blocked, ".well-known", "geul"), 0o755)
	os.WriteFile(filepath.Join(blocked, ".well-known", "geul", "_static"), nil, 0o644)
	if _, err := Tree(root, blocked); err == nil {
		t.Fatal("blocked scripts: want error")
	}
}
