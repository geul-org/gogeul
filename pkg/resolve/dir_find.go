//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.Find — one document by its 공표 주소: a version, the newest, versions.geul, documents.geul, issuers.geul or h/<해시>
package resolve

import (
	"context"
	"fmt"
	"strings"
)

// Find reads the document at addr, a 공표 주소 (GEUL-ORG) relative to
// `.well-known/geul/`, which may be written in full. Every storage form (a
// tree of files, raw and zstd pack sets, locally or over HTTP) is found the
// same way, one key at a time:
//
//	issuers.geul
//	h/<내용 해시>.geul
//	<Issuer|도메인>/documents.geul
//	<Issuer|도메인>/<문서 ID|문서 이름>.geul              the newest live version
//	<Issuer|도메인>/<문서 ID|문서 이름>/<버전 ID>.geul    that version, 0 the newest
//	<Issuer|도메인>/<문서 ID|문서 이름>/versions.geul
//
// `.geul` may be left off, as in the address without a format. IDs are
// canonical; a missing document is ErrNotFound.
func (d Dir) Find(addr string) ([]byte, error) {
	p := strings.Split(strings.TrimPrefix(strings.TrimPrefix(addr, "/"), wellKnown+"/"), "/")
	switch {
	case len(p) == 1 && p[0] == issuersFile:
		return d.Fetch(context.Background(), issuersFile)
	case len(p) == 2 && p[0] == "h":
		h := strings.TrimSuffix(p[1], geulExt)
		if _, ok := parseHexID(h, 16); !ok {
			return nil, fmt.Errorf("geul/resolve: %q is not a content hash", h)
		}
		return d.Fetch(context.Background(), "h/"+h+geulExt)
	case len(p) == 2 || len(p) == 3:
		return d.findUnder(p[0], p[1:])
	}
	return nil, fmt.Errorf("geul/resolve: %q is not a document address", addr)
}
