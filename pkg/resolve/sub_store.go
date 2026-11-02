//ff:type feature=resolve type=resolver
//ff:what subStore — the pack objects of one set: `<name>/…` of a place's objects, the place itself for the top set
package resolve

import "github.com/geul-org/gogeul/pkg/pack"

// subStore narrows base to the set name ("" is the set at the top).
type subStore struct {
	base pack.Store
	name string
}
