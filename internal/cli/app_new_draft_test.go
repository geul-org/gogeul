//ff:func feature=cli type=command control=sequence
//ff:what TestAppNewDraft — GEULISSUER fills the Issuer; a published document, an existing file without --force, a bad address and a broken key folder fail; --force replaces
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppNewDraft(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	h.env["GEULISSUER"] = "me"
	h.must("new", "0x0000000000000008", "-o", filepath.Join(h.data, "i.geul"))
	if _, err := h.run("", "new", addr, "-o", filepath.Join(h.data, "x.geul")); err == nil {
		t.Fatal("published: want error")
	}
	if _, err := h.run("", "new", "me/0x0000000000000007", "-o", draft); err == nil {
		t.Fatal("existing file: want error")
	}
	h.must("new", "me/0x0000000000000007", "-o", draft, "--force")
	if _, err := h.run("", "new", "nobody/0x0000000000000007", "-o", filepath.Join(h.data, "y.geul")); err == nil {
		t.Fatal("bad address: want error")
	}
	if _, err := h.run("", "new", "0x00000000/0x07", "-o", filepath.Join(h.data, "z.geul")); err == nil {
		t.Fatal("bad document ID: want error")
	}
	os.WriteFile(filepath.Join(h.conf, "geul", "keys", "broken.pem"), []byte("x"), 0o600)
	if _, err := h.run("", "new", "0x00000000/0x0000000000000007", "-o", filepath.Join(h.data, "w.geul")); err == nil {
		t.Fatal("broken key: want error")
	}
}
