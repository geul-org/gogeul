//ff:func feature=cli type=helper control=sequence
//ff:what TestAppResolver — the data folder comes first, then GEULPACK, then GEULPROXY; no data folder leaves the list as set; bad GEULPROXY or GEULPACK fail
package cli

import (
	"io"
	"testing"
)

func TestAppResolver(t *testing.T) {
	h := newHarness(t)
	h.doc()
	h.packed()
	set := h.data + "/.well-known/geul/0x00000000"
	h.env["GEULPROXY"] = "file:///srv,off"
	h.env["GEULPACK"] = set
	r, err := h.app(io.Discard).resolver()
	if err != nil {
		t.Fatal(err)
	}
	if p := r.Config.Proxy; len(p) != 4 || p[0].Name != "file://"+h.data || p[1].Name != "pack+"+set || p[2].Name != "file:///srv" || !p[3].Off {
		t.Fatalf("%+v", p)
	}
	h.env["GEULPACK"] = ""
	h.cwd = t.TempDir()
	h.env["GEULDATA"] = ""
	if r, err = h.app(io.Discard).resolver(); err != nil || len(r.Config.Proxy) != 2 {
		t.Fatal(r, err)
	}
	h.env["GEULPACK"] = t.TempDir()
	if _, err := h.app(io.Discard).resolver(); err == nil {
		t.Fatal("bad GEULPACK: want error")
	}
	h.env["GEULPACK"] = ""
	h.env["GEULPROXY"] = "http://plain"
	if _, err := h.app(io.Discard).resolver(); err == nil {
		t.Fatal("bad GEULPROXY: want error")
	}
}
