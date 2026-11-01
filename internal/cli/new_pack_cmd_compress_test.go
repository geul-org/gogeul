//ff:func feature=cli type=command control=sequence
//ff:what TestNewPackCmdCompress — --compress and GEULCOMPRESS pick zstd, the flag wins, an unset setting keeps the set's codec on later writes, a bad GEULCOMPRESS fails a pack write
package cli

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestNewPackCmdCompress(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	os.WriteFile(filepath.Join(h.data, ".well-known", "geul", "notes.txt"), []byte(strings.Repeat("a compressible note\n", 100)), 0o644)
	zst := regexp.MustCompile(`\([1-9][0-9]* zstd\)`).MatchString
	if out := h.must("pack", "build", "--compress", "zstd:3", "-o", filepath.Join(t.TempDir(), "a")); !zst(out) {
		t.Fatal(out)
	}
	h.env["GEULCOMPRESS"] = "zstd"
	if out := h.must("pack", "build", "-o", filepath.Join(t.TempDir(), "b")); !zst(out) {
		t.Fatal(out)
	}
	if out := h.must("pack", "build", "--compress", "none", "-o", filepath.Join(t.TempDir(), "c")); zst(out) {
		t.Fatal("flag wins", out)
	}
	delete(h.env, "GEULCOMPRESS")
	out := filepath.Join(t.TempDir(), "p")
	h.must("pack", "build", "--compress", "zstd", "-o", out)
	h.data, h.cwd = out, out
	delete(h.env, "GEULSTORE")
	h.must("add", draft, "ALIAS", "Q42", `"D. Adams"@en`)
	h.must("publish", draft)
	if log := h.must("show", "versions", addr); !strings.Contains(log, "\n3\t") {
		t.Fatal(log)
	}
	if v := h.must("pack", "compact"); !zst(v) {
		t.Fatal("follows zstd", v)
	}
	if v := h.must("pack", "compact", "--compress", "none"); zst(v) {
		t.Fatal(v)
	}
	h.must("add", draft, "ALIAS", "Q1", `"x"@en`)
	h.env["GEULCOMPRESS"] = "bogus"
	if _, err := h.run("", "publish", draft); err == nil {
		t.Fatal("GEULCOMPRESS=bogus on a pack write")
	}
}
