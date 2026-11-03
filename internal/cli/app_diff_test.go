//ff:func feature=cli type=command control=sequence
//ff:what TestAppDiff — text and JSON changes; a missing version and a document that is not all Edges fail
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppDiff(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	if out := h.must("diff", addr+"@1", "@2", "--format", "json"); !strings.Contains(out, `"op": "~"`) {
		t.Fatal(out)
	}
	if out := h.must("diff", addr+"@1", "@2"); !strings.HasPrefix(out, "~ #2") {
		t.Fatal(out)
	}
	if _, err := h.run("", "diff", addr+"@1", "@9"); err == nil {
		t.Fatal("missing: want error")
	}
	odd, _ := core.EncodeDocument(nil, []core.Record{core.EncodeUintValue(2)})
	f := filepath.Join(h.data, "odd.geul")
	os.WriteFile(f, odd, 0o644)
	if _, err := h.run("", "diff", f, f); err == nil {
		t.Fatal("not all Edges: want error")
	}
}
