//ff:func feature=build type=command control=sequence
//ff:what TestBakerDoc — bytes copied and the page written without .geul; a missing source, a blocked copy and a document that does not decode fail
package build

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestBakerDoc(t *testing.T) {
	root := fixture(t)
	src := filepath.Join(root, ".well-known", "geul")
	out := t.TempDir()
	os.Mkdir(filepath.Join(out, "h"), 0o755)
	b := &baker{tree: resolve.Dir{Root: root}.Tree(), src: src, out: out}
	if err := b.doc("h/a.geul"); err != nil {
		t.Fatal(err)
	}
	if page, err := os.ReadFile(filepath.Join(out, "h", "a")); err != nil || !strings.Contains(string(page), "browseBtn") {
		t.Fatalf("page: %v", err)
	}
	if err := b.doc("h/none.geul"); err == nil {
		t.Fatal("missing: want error")
	}
	if err := b.doc("0x00000000/x/0x00000001.geul"); err == nil {
		t.Fatal("no output directory: want error")
	}
	os.WriteFile(filepath.Join(src, "h", "bad.geul"), []byte{1, 2}, 0o644)
	if err := b.doc("h/bad.geul"); err == nil || !strings.HasPrefix(err.Error(), "h/bad.geul: ") {
		t.Fatalf("bad: %v", err)
	}
}
