//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Lookup — a document address `<Issuer>/<문서>[@<버전>]` by ID or name to its Key
package resolve

import (
	"fmt"
	"strings"
)

// Lookup resolves s, `<Issuer>/<문서>[@<버전>]`, to a Key. The Issuer is a
// Issuer ID (8, 16 or 64 hex digits) or a name of the Issuer list (GEUL-ORG Issuer 이름), the
// document a canonical 4-word ID or a name of the Issuer's document list
// (GEUL-ORG 문서 이름), and the version decimal or 2-word hex; no version is 0
// (newest).
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
	k.Doc, err = d.docOfName(k.Issuer, doc)
	return k, err
}
