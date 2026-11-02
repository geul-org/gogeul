//ff:func feature=pack type=resolver control=sequence
//ff:what TestHTTPStoreReadRange — one Range request via http.FileServer; empty ranges skip the network; a server ignoring Range, short bodies and missing objects fail
package pack

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHTTPStoreReadRange(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "f"), []byte("0123456789"), 0o644)
	c := &counter{h: http.FileServer(http.Dir(dir))}
	srv := httptest.NewServer(c)
	defer srv.Close()
	s, ctx := HTTPStore{Base: srv.URL}, context.Background()
	if b, err := s.ReadRange(ctx, "f", 2, 3); err != nil || string(b) != "234" || c.ranges != 1 {
		t.Fatal(err, c.ranges)
	}
	if b, err := s.ReadRange(ctx, "f", 2, 0); err != nil || len(b) != 0 || c.ranges != 1 {
		t.Fatal(err)
	}
	if _, err := s.ReadRange(ctx, "none", 0, 1); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := s.ReadRange(ctx, "f", 9, 3); err == nil {
		t.Fatal("past end")
	}
	whole := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("0123456789")) }))
	defer whole.Close()
	if _, err := (HTTPStore{Base: whole.URL}).ReadRange(ctx, "f", 2, 3); !errors.Is(err, ErrRange) {
		t.Fatal(err)
	}
	short := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Range", "bytes 2-4/10")
		w.WriteHeader(206)
		w.Write([]byte("23"))
	}))
	defer short.Close()
	if _, err := (HTTPStore{Base: short.URL}).ReadRange(ctx, "f", 2, 3); !errors.Is(err, ErrRange) {
		t.Fatal(err)
	}
}
