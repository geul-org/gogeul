//ff:func feature=pack type=helper control=selection
//ff:what Reader.hashOwner — the Issuer an H record is addressed under: the set's one Issuer, else the DOC-ID Issuer of its bytes, else the first Issuer
package pack

import "github.com/geul-org/gogeul/pkg/core"

// hashOwner names the Issuer whose `hashes/` serves the document b (nil when
// the bytes are not at hand). The H layer is keyed by content hash only, so
// in a set of one Issuer that Issuer owns every H record; in a set of several
// the document's own DOC-ID (DOC-POINTER or CODEC-POINTER) tells, and a
// document without one (a version list) falls to the first Issuer of the
// table, or the GEUL Issuer when the table is empty.
func (r *Reader) hashOwner(b []byte) string {
	switch {
	case len(r.m.issuers) == 1:
		return r.m.issuers[0]
	case b != nil:
		if is, ok := docIssuer(b); ok {
			return core.FormatID(is)
		}
	}
	if len(r.m.issuers) > 0 {
		return r.m.issuers[0]
	}
	return "0x00000000"
}
