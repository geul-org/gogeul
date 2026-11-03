//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestAppBlameVersions — a bad address, no data folder, an unpublished document, a version that does not read and no live version fail
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestAppBlameVersions(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	if out := h.must("blame", addr); out == "" {
		t.Fatal("want lines")
	}
	for _, a := range []string{"x", "me/0x0000000000000009"} {
		if _, err := h.run("", "blame", a); err == nil {
			t.Errorf("%s: want error", a)
		}
	}
	k, _ := h.key(addr)
	dir := filepath.Join(h.data, ".well-known", "geul", core.FormatID(k.Issuer), "docs", core.FormatID(k.Doc))
	os.WriteFile(filepath.Join(dir, "0x00000002.geul"), []byte{0, 5}, 0o644)
	if _, err := h.run("", "blame", addr); err == nil {
		t.Fatal("unreadable version: want error")
	}
	d := resolve.Dir{Root: h.data}
	d.Tombstone(k.Issuer, k.Doc, 1, 1)
	if _, err := h.run("", "blame", addr+"@1"); err == nil {
		t.Fatal("no live version: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "blame", "0x00000000/0x0000000000000001"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
