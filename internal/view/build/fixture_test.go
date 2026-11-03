//ff:func feature=build type=helper control=sequence
//ff:what fixture — test fixture: a tree with an empty document under h/, a document under 0x00000000/x/, and a file that is not .geul
package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	top := filepath.Join(root, ".well-known", "geul")
	os.MkdirAll(filepath.Join(top, "h"), 0o755)
	os.MkdirAll(filepath.Join(top, "0x00000000", "x"), 0o755)
	doc := core.PutWords(core.WrapDocument([]core.Word{0x1C40}))
	os.WriteFile(filepath.Join(top, "h", "a.geul"), doc, 0o644)
	os.WriteFile(filepath.Join(top, "0x00000000", "x", "0x00000001.geul"), doc, 0o644)
	os.WriteFile(filepath.Join(top, "h", "notes.txt"), []byte("x"), 0o644)
	return root
}
