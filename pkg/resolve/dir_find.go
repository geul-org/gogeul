//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.Find — one document by its 공표 주소 (the sixteen forms), the Issuer also as a name of the Issuer list and the document as a name of the document list
package resolve

import (
	"context"
	"strings"
)

// Find reads the document at addr, a 공표 주소 (§3.13) relative to
// `.well-known/geul/`, which may be written in full. Every storage form (a
// tree of files, raw and zstd pack sets, locally or over HTTP) is found the
// same way, one key at a time:
//
//	issuers/<latest|<버전 ID>|versions>.geul
//	<Issuer>/codecs/[<코덱 ID>/]<latest|<버전 ID>|versions>.geul
//	<Issuer>/docs/[<문서 ID>/]<latest|<버전 ID>|versions>.geul
//	<Issuer>/hashes/<내용 해시>.geul
//
// For text viewing the Issuer may also be a name of the Issuer list and the
// document a name of the Issuer's document list, and `.geul` may be left off.
// A codec or document version is read as a snapshot (At: a stored diff is
// applied), latest.geul being the newest live version. A missing document is
// ErrNotFound.
func (d Dir) Find(addr string) ([]byte, error) {
	p := strings.TrimPrefix(strings.TrimPrefix(addr, "/"), wellKnown+"/")
	if !strings.HasSuffix(p, geulExt) {
		p += geulExt
	}
	a, err := addrOfPath(p, d.issuerOfName, d.docOfName)
	switch {
	case err != nil:
		return nil, err
	case (a.Kind == AddrDoc || a.Kind == AddrCodec) && !a.Versions:
		return d.At(a.Issuer, a.ID, a.Version)
	default:
		return FetchAddr(context.Background(), d, a)
	}
}
