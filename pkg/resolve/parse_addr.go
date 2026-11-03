//ff:func feature=resolve type=helper control=sequence
//ff:what ParseAddr — a 공표 주소 (file://, https:// or a path under .well-known/geul/) as an Addr; the ten forms only, else ErrNotFound
package resolve

import (
	"fmt"
	"strings"
)

// ParseAddr reads s, one of the ten 공표 주소 forms (§3.13), as an Addr:
// `file://<dir>/.well-known/geul/<path>`, `https://<host>/.well-known/geul/<path>`
// or a bare `<path>` (optionally starting with `.well-known/geul/`). IDs only:
// the Issuer as 8, 16 or 64 hex digits (the same Issuer), the codec ID 8 and
// the document ID 16 digits, the version 8 digits and not 0. Any other form is
// ErrNotFound — the address does not exist (geul get, HTTP 404).
func ParseAddr(s string) (Addr, error) {
	base, path := "", s
	if strings.HasPrefix(s, "file://") || strings.HasPrefix(s, "https://") {
		b, p, ok := strings.Cut(s, "/"+wellKnown+"/")
		if !ok {
			return Addr{}, fmt.Errorf("%w: %q has no /%s/", ErrNotFound, s, wellKnown)
		}
		base, path = b, p
	}
	path = strings.TrimPrefix(path, wellKnown+"/")
	a, err := addrOfPath(path, issuerOfID)
	a.Base = base
	return a, err
}
