//ff:func feature=notation type=helper control=selection
//ff:what fakeLookup — test fixture: a Lookup that knows geul.org/0x0000000000000806 and the 16-word Issuer "self", else reads IDs
package notation

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// fakeLookup resolves two names and passes IDs to resolve.Dir{}.Lookup.
func fakeLookup(addr string) (resolve.Key, error) {
	switch addr {
	case "geul.org/0x0000000000000806":
		return resolve.Key{Issuer: []core.Word{0, 0}, Doc: []core.Word{0, 0, 0, 0x806}}, nil
	case "self/0x0000000000000000", "self/0x0000000000000001":
		is := make([]core.Word, 16)
		is[0] = 1
		return resolve.Key{Issuer: is, Doc: []core.Word{0, 0, 0, 1}}, nil
	default:
		return resolve.Dir{}.Lookup(addr)
	}
}
