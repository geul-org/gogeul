//ff:func feature=server type=handler control=iteration dimension=1
//ff:what TestServePage — page with CSP, 404 off root, 500 for a missing file, 422 for a non-document, 500 when writing fails
package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestServePage(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "a.geul")
	bad := filepath.Join(dir, "b.geul")
	os.WriteFile(good, core.PutWords(core.WrapDocument([]core.Word{0x1C40})), 0o644)
	os.WriteFile(bad, []byte{0x18, 0x0F}, 0o644)
	get := func(w http.ResponseWriter, path, url string) {
		servePage(w, httptest.NewRequest(http.MethodGet, url, nil), path)
	}
	rec := httptest.NewRecorder()
	get(rec, good, "/")
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "<title>a.geul · geul</title>") || rec.Header().Get("Content-Security-Policy") != contentPolicy {
		t.Fatalf("page: %d %s", rec.Code, rec.Header())
	}
	for _, c := range []struct {
		path, url string
		code      int
	}{{good, "/x", 404}, {filepath.Join(dir, "none.geul"), "/", 500}, {bad, "/", 422}} {
		rec := httptest.NewRecorder()
		get(rec, c.path, c.url)
		if rec.Code != c.code {
			t.Fatalf("%s %s: %d want %d", c.path, c.url, rec.Code, c.code)
		}
	}
	bw := brokenWriter{httptest.NewRecorder()}
	get(bw, good, "/")
	if bw.Code != 500 {
		t.Fatalf("broken writer: %d", bw.Code)
	}
}
