//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestAppLastTwo — the newest against the one before; one version is too few; deleted versions are skipped; bad addresses and no data folder fail
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppLastTwo(t *testing.T) {
	h := newHarness(t)
	h.must("issuer", "new", "--name", "me")
	f := filepath.Join(h.data, "n.geul")
	h.must("new", "me/0x0000000000000005", "-o", f)
	h.must("add", f, "LABEL", "Q1", `"a"`)
	h.must("publish", f)
	if _, err := h.run("", "diff", "me/0x0000000000000005"); err == nil {
		t.Fatal("one version: want error")
	}
	h.must("add", f, "ALIAS", "Q1", `"b"`)
	h.must("publish", f)
	if out := h.must("diff", "me/0x0000000000000005"); out != "+ #2 ALIAS Q1 \"b\"\n" {
		t.Fatalf("%q", out)
	}
	k, _ := h.key("me/0x0000000000000005")
	os.WriteFile(filepath.Join(h.data, ".well-known", "geul", core.FormatID(k.Issuer), "docs", core.FormatID(k.Doc), "0x00000001.geul"), []byte{0, 5}, 0o644)
	if _, err := h.run("", "diff", "me/0x0000000000000005"); err == nil {
		t.Fatal("unreadable older version: want error")
	}
	for _, a := range []string{"x", "me/0x0000000000000009"} {
		if _, err := h.run("", "diff", a); err == nil {
			t.Errorf("%s: want error", a)
		}
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "diff", "0x00000000/0x0000000000000001"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
