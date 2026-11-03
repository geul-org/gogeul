//ff:func feature=server type=handler control=sequence
//ff:what serveByName — redirect an Issuer name path to its ID path; 404 when the first segment is not a known Issuer name
package server

import (
	"net/http"
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// serveByName answers a tree path that does not exist by reading its first
// segment as an Issuer name (namePath; documents have no names). A known name
// redirects (302) to the same path under the Issuer ID, keeping a trailing /
// and the query. Anything else is 404.
func serveByName(w http.ResponseWriter, r *http.Request, t resolve.Store, p string) {
	to, ok := namePath(t, p)
	if !ok {
		http.NotFound(w, r)
		return
	}
	if strings.HasSuffix(r.URL.Path, "/") {
		to += "/"
	}
	if r.URL.RawQuery != "" {
		to += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, to, http.StatusFound)
}
