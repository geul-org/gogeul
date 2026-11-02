//ff:func feature=pack type=resolver control=sequence
//ff:what TestHTTPStoreReadAll — 200 body; other statuses and errors fail
package pack

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPStoreReadAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/boom" {
			w.WriteHeader(500)
			return
		}
		if r.URL.Path == "/none" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("ok"))
	}))
	defer srv.Close()
	s, ctx := HTTPStore{Base: srv.URL}, context.Background()
	if b, err := s.ReadAll(ctx, "latest.txt"); err != nil || string(b) != "ok" {
		t.Fatal(err)
	}
	if _, err := s.ReadAll(ctx, "boom"); err == nil {
		t.Fatal("500")
	}
	if _, err := s.ReadAll(ctx, "none"); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
