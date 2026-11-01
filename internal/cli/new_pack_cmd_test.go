//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestNewPackCmd — build, verify, compact, unpack through the command line, and pack writes from publish; bad --compress values and missing data folders fail
package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNewPackCmd(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	out := filepath.Join(t.TempDir(), "p")
	h.must("pack", "build", "-o", out)
	h.data, h.cwd = out, out
	delete(h.env, "GEULSTORE")
	if v := h.must("pack", "verify"); strings.Count(v, " ok\n") != 2 {
		t.Fatal(v)
	}
	h.must("add", draft, "ALIAS", "Q42", `"D. Adams"@en`)
	h.must("publish", draft)
	if log := h.must("show", "versions", addr); !strings.Contains(log, "\n3\t") {
		t.Fatal(log)
	}
	h.must("pack", "compact")
	h.must("pack", "unpack", "-o", filepath.Join(t.TempDir(), "u"))
	h.must("check", "--archive")
	for _, sub := range []string{"build --compress gzip -o x", "compact --compress zstd:99"} {
		if _, err := h.run("", append([]string{"pack"}, strings.Fields(sub)...)...); err == nil {
			t.Fatal(sub, ": bad codec: want error")
		}
	}
	h.cwd = t.TempDir()
	for _, sub := range []string{"build -o x", "unpack -o x", "verify", "compact"} {
		if _, err := h.run("", append([]string{"pack"}, strings.Fields(sub)...)...); err == nil {
			t.Fatal(sub, ": no data folder: want error")
		}
	}
}
