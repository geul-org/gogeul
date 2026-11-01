//ff:func feature=pack type=helper control=sequence
//ff:what ballast — test fixture: one more small file in an archive tree, so the set has many index blocks
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

// ballast writes path under root's `.well-known/geul/`.
func ballast(t *testing.T, root, path string) {
	t.Helper()
	f := filepath.Join(root, ".well-known", "geul", filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(f), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(f, []byte(path), 0o644); err != nil {
		t.Fatal(err)
	}
}
