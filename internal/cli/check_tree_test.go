//ff:func feature=cli type=helper control=sequence
//ff:what TestCheckTree — the same findings, under the same file names, for a directory tree and its pack sets; walk errors fail
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestCheckTree(t *testing.T) {
	h := newHarness(t)
	h.doc()
	os.MkdirAll(filepath.Join(h.data, ".well-known", "geul", "issuers"), 0o755)
	os.WriteFile(filepath.Join(h.data, ".well-known", "geul", "issuers", "0x00000009.geul"), []byte{1, 2}, 0o644)
	dir, err := checkTree(resolve.Dir{Root: h.data}.Tree(), "base")
	if err != nil || len(dir) == 0 || !strings.HasPrefix(dir[0].Path, "base") {
		t.Fatal(dir, err)
	}
	h.packed()
	tr, _, _ := openTree(h.data)
	packed, err := checkTree(tr, "base")
	if err != nil || fmt.Sprint(packed) != fmt.Sprint(dir) {
		t.Fatalf("%v\n%v %v", dir, packed, err)
	}
	if _, err := checkTree(failTree{fstest.MapFS{"a.geul": {}}}, "b"); err == nil {
		t.Fatal("read error")
	}
}
