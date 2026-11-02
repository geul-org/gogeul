//ff:func feature=resolve type=resolver control=sequence
//ff:what remoteSite — the site of an `https://` base: pack objects by HTTP Range, loose files by HTTP GET
package resolve

import (
	"net/http"
	"strings"

	"github.com/geul-org/gogeul/pkg/pack"
)

// remoteSite reads `<base>/.well-known/geul/`. Client nil uses http.DefaultClient.
func remoteSite(base string, c *http.Client) *site {
	wk := strings.TrimSuffix(base, "/") + "/" + wellKnown
	return &site{objects: pack.HTTPStore{Base: wk, Client: c}, loose: HTTP{Base: base, Client: c}, sets: map[string]*pack.Reader{}}
}
