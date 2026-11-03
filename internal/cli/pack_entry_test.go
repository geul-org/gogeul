//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestPackEntry — an http(s) set base, a local set directory and a packed data folder become read-only sources; other folders and dead servers fail
package cli

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestPackEntry(t *testing.T) {
	h := newHarness(t)
	h.doc()
	h.packed()
	set := filepath.Join(h.data, ".well-known", "geul", "0x00000000")
	srv := httptest.NewServer(http.FileServer(http.Dir(set)))
	defer srv.Close()
	for _, item := range []string{srv.URL, set, h.data} {
		e, err := packEntry(item)
		if err != nil || e.Name != "pack+"+item {
			t.Fatal(item, err)
		}
		if d, ok := e.Source.(resolve.Dir); !ok || d.Store == nil {
			t.Fatal(item, "not a pack source")
		}
	}
	if _, err := packEntry(t.TempDir()); err == nil {
		t.Fatal("plain folder")
	}
	srv.Close()
	if _, err := packEntry(srv.URL); err == nil {
		t.Fatal("dead server")
	}
}
