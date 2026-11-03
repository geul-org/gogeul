//ff:func feature=cli type=command control=sequence
//ff:what TestAppShow — a draft file, JSON; a missing document and a broken file fail
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppShow(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	if out := h.must("show", draft); !strings.Contains(out, "#3 ALIAS Q42 \"DNA\"@en") {
		t.Fatal(out)
	}
	if out := h.must("show", draft, "--format", "json"); !strings.Contains(out, `"body"`) {
		t.Fatal(out)
	}
	if _, err := h.run("", "show", "me/docs/0x0000000000000009/latest"); err == nil {
		t.Fatal("missing: want error")
	}
	bad := filepath.Join(h.data, "bad.geul")
	os.WriteFile(bad, []byte{1}, 0o644)
	if _, err := h.run("", "show", bad); err == nil {
		t.Fatal("broken: want error")
	}
}
