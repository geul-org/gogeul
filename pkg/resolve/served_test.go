//ff:func feature=resolve type=helper control=sequence
//ff:what served — test fixture: a data folder behind an HTTP server that answers Range requests, read as a remote site
package resolve

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// served returns a remote site over root and the server's meter.
func served(t *testing.T, root string) (*site, *requestMeter) {
	t.Helper()
	m := &requestMeter{next: http.FileServer(http.Dir(root))}
	srv := httptest.NewServer(m)
	t.Cleanup(srv.Close)
	return remoteSite(srv.URL, srv.Client()), m
}
