//ff:func feature=server type=handler control=iteration dimension=1
//ff:what TestServeByName — known Issuer and document names redirect to their ID path (trailing / and query kept); unknown names, a missing or broken issuers.geul are 404
package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestServeByName(t *testing.T) {
	root := namesFixture(t)
	cases := map[string]string{
		"/.well-known/geul/geul.org/docs/0xffff000000000002/0x00000001.geul": "/.well-known/geul/0x00000000/docs/0xffff000000000002/0x00000001.geul",
		"/.well-known/geul/geul.org/?json":                                   "/.well-known/geul/0x00000000/?json",
		"/.well-known/geul/geul.org":                                         "/.well-known/geul/0x00000000",
		"/.well-known/geul/geul.org/sample-b/0x00000001.geul":                "/.well-known/geul/0x00000000/docs/0xffff000000000002/0x00000001.geul",
		"/.well-known/geul/0x00000000/sample-b/":                             "/.well-known/geul/0x00000000/docs/0xffff000000000002/",
	}
	for u, want := range cases {
		rec := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, u, nil)
		serveTree(rec, r, treeOf(root))
		if rec.Code != http.StatusFound || rec.Header().Get("Location") != want {
			t.Fatalf("%s: %d %q", u, rec.Code, rec.Header().Get("Location"))
		}
	}
	rec := httptest.NewRecorder()
	serveTree(rec, httptest.NewRequest(http.MethodGet, "/.well-known/geul/nope.org/x", nil), treeOf(root))
	if rec.Code != 404 {
		t.Fatalf("unknown: %d", rec.Code)
	}
	issuers := filepath.Join(root, ".well-known", "geul", "issuers/0x00000001.geul")
	os.WriteFile(issuers, []byte{1}, 0o644)
	rec = httptest.NewRecorder()
	serveTree(rec, httptest.NewRequest(http.MethodGet, "/.well-known/geul/geul.org/x", nil), treeOf(root))
	if rec.Code != 404 {
		t.Fatalf("broken: %d", rec.Code)
	}
	os.Remove(issuers)
	rec = httptest.NewRecorder()
	serveTree(rec, httptest.NewRequest(http.MethodGet, "/.well-known/geul/geul.org/x", nil), treeOf(root))
	if rec.Code != 404 {
		t.Fatalf("missing: %d", rec.Code)
	}
}
