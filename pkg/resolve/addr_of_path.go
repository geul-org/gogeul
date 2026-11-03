//ff:func feature=resolve type=helper control=selection
//ff:what addrOfPath — read the segments of a 공표 주소 path into an Addr; anything outside the ten forms is ErrNotFound
package resolve

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// addrOfPath reads p, a path relative to `.well-known/geul/`, as one of the
// ten forms (§3.13). issuerOf reads the Issuer segment (ParseAddr takes IDs
// only; Dir.Find also takes Issuer names). Documents and codecs are named by
// ID only. Any other spelling, `<Issuer>/docs/` and `<Issuer>/codecs/`
// themselves included (no lists, §1.1 101), is ErrNotFound, as the address
// does not exist.
func addrOfPath(p string, issuerOf func(string) ([]core.Word, error)) (Addr, error) {
	seg := strings.Split(p, "/")
	bad := fmt.Errorf("%w: %q is not a 공표 주소", ErrNotFound, p)
	if len(seg) == 2 && seg[0] == issuersDir {
		return addrTail(Addr{Kind: AddrIssuers}, seg[1], bad)
	}
	if len(seg) < 3 {
		return Addr{}, bad
	}
	issuer, err := issuerOf(seg[0])
	if err != nil {
		return Addr{}, err
	}
	a := Addr{Issuer: issuer}
	switch {
	case seg[1] == hashesDir && len(seg) == 3:
		return addrHash(a, seg[2], bad)
	case seg[1] == codecsDir && len(seg) == 4:
		return addrCodec(a, seg[2], seg[3], bad)
	case seg[1] == docsDir && len(seg) == 4:
		return addrDoc(a, seg[2], seg[3], bad)
	default:
		return Addr{}, bad
	}
}
