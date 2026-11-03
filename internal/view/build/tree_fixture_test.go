//ff:func feature=build type=helper control=sequence
//ff:what treeFixture — test fixture: a temp root with .well-known/geul/h/a.geul, a bad .geul and a versions file
package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func treeFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	h := filepath.Join(root, ".well-known", "geul", "h")
	if err := os.MkdirAll(h, 0o755); err != nil {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(h, "a.geul"), core.PutWords(core.WrapDocument([]core.Word{0x1C40})), 0o644)
	os.WriteFile(filepath.Join(h, "bad.geul"), []byte{1, 2}, 0o644)
	os.WriteFile(filepath.Join(h, "versions"), []byte("x"), 0o644)
	os.WriteFile(filepath.Join(root, "outside.geul"), core.PutWords(core.WrapDocument([]core.Word{0x1C40})), 0o644)
	return root
}
