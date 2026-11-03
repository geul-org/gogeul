//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Lookup — a document address `<Issuer>/<문서>[@<버전>]` to its Key: the Issuer by ID or name, the document by ID
package resolve

import (
	"fmt"
	"strings"
)

// Lookup resolves s, `<Issuer>/<문서>[@<버전>]`, to a Key. The Issuer is a
// Issuer ID (8, 16 or 64 hex digits) or a name of the Issuer list (GEUL-ORG Issuer 이름), the
// document a canonical 4-word document ID or 2-word codec ID (lookupDocID; no
// document names, §1.1 101), and the version decimal or 2-word hex; no
// version is 0 (newest).
func (d Dir) Lookup(s string) (Key, error) {
	addr, ver, hasVer := strings.Cut(s, "@")
	var k Key
	if hasVer {
		v, err := parseVersion(ver)
		if err != nil {
			return Key{}, err
		}
		k.Version = v
	}
	is, doc, ok := strings.Cut(addr, "/")
	if !ok {
		return Key{}, fmt.Errorf("geul/resolve: %q is not <Issuer>/<문서>", s)
	}
	issuer, err := d.issuerOfName(is)
	if err != nil {
		return Key{}, err
	}
	k.Issuer = issuer
	k.Doc, err = lookupDocID(k.Issuer, doc)
	return k, err
}
