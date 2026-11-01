//ff:func feature=cli type=command control=sequence
//ff:what TestAppPublish — --amend leaves one version; the same content again is noted, not published; a missing file, a document breaking the spec, no data folder, no DOC-ID and a taken name fail
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppPublish(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	if out := h.must("publish", draft, "--amend"); !strings.HasPrefix(out, "1 0x") {
		t.Fatal(out)
	}
	if out := h.must("show", "versions", addr); strings.Count(out, "\n") != 2 {
		t.Fatal(out)
	}
	if out := h.must("publish", draft); !strings.HasPrefix(out, "1 0x") || !strings.Contains(out, "already published") {
		t.Fatal(out)
	}
	if _, err := h.run("", "publish", filepath.Join(h.data, "none")); err == nil {
		t.Fatal("missing: want error")
	}
	bad := filepath.Join(h.data, "bad.geul")
	os.WriteFile(bad, []byte{0, 5}, 0o644)
	if _, err := h.run("", "publish", bad); err == nil {
		t.Fatal("breaks the spec: want error")
	}
	noID, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, nil)
	g := filepath.Join(h.data, "noid.geul")
	os.WriteFile(g, noID, 0o644)
	if _, err := h.run("", "publish", g); err == nil {
		t.Fatal("no DOC-ID: want error")
	}
	h.must("name", "me/0x0000000000000001", "taken")
	f := filepath.Join(h.data, "two.geul")
	h.must("new", "me/0x0000000000000002", "-o", f)
	if _, err := h.run("", "publish", f, "--name", "taken"); err == nil {
		t.Fatal("taken name: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "publish", draft); err == nil {
		t.Fatal("no data folder: want error")
	}
}
