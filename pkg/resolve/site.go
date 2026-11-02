//ff:type feature=resolve type=resolver
//ff:what site — one place's `.well-known/geul/` whose form is found out while reading: pack sets opened on first use, then loose files
package resolve

import (
	"sync"

	"github.com/geul-org/gogeul/pkg/pack"
)

// site reads a GEULPROXY place without being told how it stores documents. A
// path is looked up in the pack sets that may serve it (setsFor), then as a
// loose file, so a tree of `.geul` files, raw and zstd pack sets and any mix
// of them per Issuer read alike. objects reads pack objects under
// `.well-known/geul/`, loose reads one path; dir is the local directory, "" for
// a remote place, which has no listing. sets remembers each set name tried (a
// nil Reader is "no set there") for the life of the site.
type site struct {
	objects pack.Store
	loose   Source
	dir     string

	mu   sync.Mutex
	sets map[string]*pack.Reader
}
