//ff:type feature=notation type=model
//ff:what Context — what reading a line needs: codec specs for field layouts, a Lookup for names, a map for #labels
package notation

import "github.com/geul-org/gogeul/pkg/codecspec"

// Context configures reading. Spec gives the fields of declared codecs (a
// codec without one is read token by token). Lookup resolves Issuer names.
// Local, when set, maps a `#label` written in a token (the text after #) to
// the LOCAL ID to encode (labels of a batch, geul add -i); a label it does not
// know is read as the decimal Edge Id it spells.
type Context struct {
	Spec   codecspec.Spec
	Lookup Lookup
	Local  func(label string) (uint64, bool)
}
