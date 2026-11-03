//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what fixture — a writable copy of testdata/archive (4 documents, their hashes/ copies, one codec version and the Issuer list issuers/)
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

// fixture copies the test archive into a temp root.
func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	paths, err := listTree("testdata/archive")
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range paths {
		dst := filepath.Join(root, ".well-known", "geul", filepath.FromSlash(p))
		os.MkdirAll(filepath.Dir(dst), 0o755)
		if err := os.WriteFile(dst, fixtureFile(t, p), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}
