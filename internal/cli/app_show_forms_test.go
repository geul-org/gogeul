//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppShowForms — latest, a given version, a local file; a missing document and version fail; a version number and JSON; ./versions is a file
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppShowForms(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	if out := h.must("show", addr); !strings.Contains(out, "#2 TRIPLE P31 Q42 Q36180 normal\n") || !strings.Contains(out, "ALIAS") {
		t.Fatal(out)
	}
	if out := h.must("show", addr+"@1"); !strings.Contains(out, "#2 TRIPLE P31 Q42 Q5 normal\n") || strings.Contains(out, "ALIAS") {
		t.Fatal(out)
	}
	if a, b := h.must("show", draft), h.must("show", addr); a != b {
		t.Fatalf("file %q\naddress %q", a, b)
	}
	for _, missing := range []string{addr + "@9", strings.TrimSuffix(addr, "1") + "9"} {
		if _, err := h.run("", "show", missing); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Fatal(missing, err)
		}
	}
	b, _ := os.ReadFile(draft)
	os.WriteFile(filepath.Join(h.data, "versions"), b, 0o644)
	t.Chdir(h.data)
	if out := h.must("show", "./versions"); out != h.must("show", draft) {
		t.Fatal(out)
	}
	if out := h.must("show", "versions"); out != h.must("show", draft) {
		t.Fatal("versions alone is an ordinary argument:", out)
	}
	if out := h.must("show", "versions", addr, "--format", "json"); !strings.Contains(out, `"next_local_id"`) {
		t.Fatal(out)
	}
	if _, err := h.run("", "show", "versions", "./versions"); err == nil {
		t.Fatal("show versions takes a document, not a file")
	}
}
