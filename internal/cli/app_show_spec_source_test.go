//ff:func feature=cli type=command control=sequence
//ff:what TestAppShowSpecSource — a spec missing from the data folder is found under 0x0000 through GEULPROXY and lays out the document
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppShowSpecSource(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	with := h.must("show", docsAddr(addr)+"/latest")
	const rel = ".well-known/geul/0x00000000/codecs/0x00000AA0"
	other := t.TempDir()
	os.MkdirAll(filepath.Dir(filepath.Join(other, rel)), 0o755)
	if err := os.Rename(filepath.Join(h.data, rel), filepath.Join(other, rel)); err != nil {
		t.Fatal(err)
	}
	if out := h.must("show", docsAddr(addr)+"/latest"); out == with {
		t.Fatal("the spec is gone from the data folder and GEULPROXY is off:", out)
	}
	h.env["GEULPROXY"] = "file://" + other
	if out := h.must("show", docsAddr(addr)+"/latest"); out != with {
		t.Fatalf("got %q\nwant %q", out, with)
	}
}
