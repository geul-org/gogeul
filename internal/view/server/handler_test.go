//ff:func feature=server type=handler control=sequence
//ff:what TestHandler — the handler serves the page of its file and the embedded scripts
package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHandler(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.geul")
	if err := os.WriteFile(path, core.PutWords(core.WrapDocument([]core.Word{0x1C40})), 0o644); err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	Handler(path).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d: %s", rec.Code, rec.Body)
	}
	rec = httptest.NewRecorder()
	Handler(path).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/.well-known/geul/_static/cytoscape.min.js", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("static: %d", rec.Code)
	}
}
