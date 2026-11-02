//ff:func feature=resolve type=helper control=sequence
//ff:what requestMeter.ServeHTTP — record the request and its bytes, then answer it
package resolve

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

// ServeHTTP answers from fail or the folder.
func (m *requestMeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	if m.fail == nil || !m.fail(rec, r) {
		m.next.ServeHTTP(rec, r)
	}
	m.mu.Lock()
	m.seen = append(m.seen, strings.TrimPrefix(r.URL.Path, "/"+wellKnown+"/")+" "+r.Header.Get("Range"))
	m.sent += int64(rec.Body.Len())
	m.mu.Unlock()
	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	w.WriteHeader(rec.Code)
	w.Write(rec.Body.Bytes())
}
