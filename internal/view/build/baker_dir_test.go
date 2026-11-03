//ff:func feature=build type=command control=iteration dimension=1
//ff:what TestBakerDir — the listing page and its JSON are written; an unmakable directory, a missing source directory and a blocked JSON fail
package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestBakerDir(t *testing.T) {
	root := fixture(t)
	out := t.TempDir()
	b := &baker{tree: resolve.Dir{Root: root}.Tree(), out: out}
	if err := b.dir("h/"); err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{"index.html", "index.json"} {
		if _, err := os.Stat(filepath.Join(out, "h", f)); err != nil {
			t.Fatalf("%s: %v", f, err)
		}
	}
	os.WriteFile(filepath.Join(out, "file"), nil, 0o644)
	if err := b.dir("file/"); err == nil {
		t.Fatal("under a file: want error")
	}
	if err := b.dir("none/"); err == nil {
		t.Fatal("missing source: want error")
	}
	os.Mkdir(filepath.Join(out, "0x00000000"), 0o755)
	os.Mkdir(filepath.Join(out, "0x00000000", "index.json"), 0o755)
	if err := b.dir("0x00000000/"); err == nil {
		t.Fatal("blocked json: want error")
	}
}
