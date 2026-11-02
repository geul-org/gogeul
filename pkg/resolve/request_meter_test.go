//ff:type feature=resolve type=model
//ff:what requestMeter — test fixture: a handler over a data folder recording each request's path and Range
package resolve

import (
	"net/http"
	"sync"
)

// requestMeter counts what a remote read costs. fail, when set, answers a
// request instead of the folder.
type requestMeter struct {
	next http.Handler
	fail func(http.ResponseWriter, *http.Request) bool

	mu   sync.Mutex
	seen []string // "<path under .well-known/geul/> <Range>"
	sent int64
}
