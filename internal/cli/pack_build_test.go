//ff:func feature=cli type=command control=sequence
//ff:what TestPackBuild — per-Issuer sets or one set; the same folder, a pack archive as source, no tree and unreadable files fail
package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackBuild(t *testing.T) {
	h := newHarness(t)
	h.doc()
	var out bytes.Buffer
	split := filepath.Join(t.TempDir(), "split")
	if err := packBuild(h.data, split, false, packOptions(), &out); err != nil || strings.Count(out.String(), "generation 1") != 2 {
		t.Fatal(err, out.String())
	}
	single := filepath.Join(t.TempDir(), "single")
	if err := packBuild(filepath.Join(h.data, ".well-known", "geul"), single, true, packOptions(), &out); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(single, ".well-known", "geul", "latest.txt")); err != nil {
		t.Fatal(err)
	}
	if err := packBuild(h.data, h.data, false, packOptions(), &out); err == nil {
		t.Fatal("same folder")
	}
	if err := packBuild(split, t.TempDir(), false, packOptions(), &out); err == nil {
		t.Fatal("already packed")
	}
	if err := packBuild(t.TempDir(), t.TempDir(), false, packOptions(), &out); err == nil {
		t.Fatal("no tree")
	}
	os.MkdirAll(filepath.Join(h.data, ".well-known", "geul", "0x00000000", "hashes"), 0o755)
	os.WriteFile(filepath.Join(h.data, ".well-known", "geul", "0x00000000", "hashes", "0x"+strings.Repeat("0", 64)+".geul"), []byte("x"), 0o644)
	if err := packBuild(h.data, filepath.Join(t.TempDir(), "bad"), true, packOptions(), &out); err == nil {
		t.Fatal("bad h/ copy")
	}
}
