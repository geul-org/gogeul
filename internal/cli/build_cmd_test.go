//ff:func feature=cli type=command control=sequence
//ff:what TestBuildCmd — a directory without a tree, a build that fails, a failing output, a broken pack archive, and a summary on success
package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildCmd(t *testing.T) {
	if err := buildCmd(t.TempDir(), t.TempDir(), &bytes.Buffer{}); err == nil || !strings.Contains(err.Error(), "no .well-known/geul") {
		t.Fatalf("no tree: %v", err)
	}
	root := t.TempDir()
	top := filepath.Join(root, ".well-known", "geul")
	os.MkdirAll(top, 0o755)
	if err := buildCmd(root, root, &bytes.Buffer{}); err == nil {
		t.Fatal("output over source: want error")
	}
	if err := buildCmd(root, t.TempDir(), failWriter{}); err == nil {
		t.Fatal("output failure: want error")
	}
	os.MkdirAll(filepath.Join(top, "0x00000000"), 0o755)
	os.WriteFile(filepath.Join(top, "0x00000000", "latest.txt"), []byte("x"), 0o644)
	if err := buildCmd(root, t.TempDir(), &bytes.Buffer{}); err == nil {
		t.Fatal("broken pack archive: want error")
	}
	os.RemoveAll(filepath.Join(top, "0x00000000"))
	var out bytes.Buffer
	dst := t.TempDir()
	if err := buildCmd(top, dst, &out); err != nil || out.String() != "geul: 0 documents, 1 directories → "+filepath.Join(dst, ".well-known", "geul")+"\n" {
		t.Fatalf("summary: %v %q", err, out.String())
	}
}
