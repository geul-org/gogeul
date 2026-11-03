//ff:func feature=cli type=helper control=sequence
//ff:what TestCheckVersions — a broken list and a version that does not read back are findings; deleted versions are skipped
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestCheckVersions(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	k, _ := h.key(addr)
	d := resolve.Dir{Root: h.data}
	d.Tombstone(k.Issuer, k.Doc, 1, 1)
	if fs := checkVersions(d, k.Issuer, k.Doc); len(fs) != 0 {
		t.Fatalf("deleted skipped: %+v", fs)
	}
	dir := filepath.Join(h.data, ".well-known", "geul", core.FormatID(k.Issuer), "docs", core.FormatID(k.Doc))
	os.WriteFile(filepath.Join(dir, "0x00000002.geul"), []byte{0, 5}, 0o644)
	if fs := checkVersions(d, k.Issuer, k.Doc); len(fs) != 1 {
		t.Fatalf("bad version: %+v", fs)
	}
	os.WriteFile(filepath.Join(dir, "versions.geul"), []byte{1}, 0o644)
	if fs := checkVersions(d, k.Issuer, k.Doc); len(fs) != 1 {
		t.Fatalf("bad list: %+v", fs)
	}
}
