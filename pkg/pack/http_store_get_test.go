//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestHTTPStoreGet — headers are sent; 404 and 410 are ErrNotFound; a bad URL and a dead server fail
package pack

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPStoreGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/b/gone":
			w.WriteHeader(http.StatusGone)
		case "/b/x":
			w.Write([]byte(r.Header.Get("X-T")))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	ctx := context.Background()
	s := HTTPStore{Base: srv.URL + "/b/"}
	resp, err := s.get(ctx, "x", "X-T", "v")
	if err != nil || resp.StatusCode != 200 {
		t.Fatal(err)
	}
	resp.Body.Close()
	for _, n := range []string{"gone", "none"} {
		if _, err := s.get(ctx, n); !errors.Is(err, ErrNotFound) {
			t.Fatal(n, err)
		}
	}
	if _, err := (HTTPStore{Base: "http://bad host"}).get(ctx, "x"); err == nil {
		t.Fatal("bad url")
	}
	dead := httptest.NewServer(http.NotFoundHandler())
	dead.Close()
	if _, err := (HTTPStore{Base: dead.URL, Client: &http.Client{}}).get(ctx, "x"); err == nil {
		t.Fatal("dead")
	}
}
