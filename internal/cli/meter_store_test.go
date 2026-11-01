//ff:type feature=cli type=helper
//ff:what meterStore — test fixture: a pack.Store that counts range reads, so a listing that walks the index shows
package cli

import "github.com/geul-org/gogeul/pkg/pack"

// meterStore counts ReadRange calls on Store.
type meterStore struct {
	pack.Store
	reads int
}
