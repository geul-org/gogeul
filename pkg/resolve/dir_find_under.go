//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.findUnder — the part of a 공표 주소 under the Issuer: documents.geul, the newest, versions.geul or a version
package resolve

import (
	"context"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// findUnder reads rest, one or two address segments, under the Issuer written
// as is, an ID or a name of `issuers.geul`. A document name goes through the
// Issuer's `documents.geul`, the newest version through `versions.geul` (At).
func (d Dir) findUnder(is string, rest []string) ([]byte, error) {
	issuer, err := d.issuerOfName(is)
	if err != nil {
		return nil, err
	}
	if rest[0] == namesFile && len(rest) == 1 {
		return d.Fetch(context.Background(), core.FormatID(issuer)+"/"+namesFile)
	}
	doc, err := d.docOfName(issuer, strings.TrimSuffix(rest[0], geulExt))
	switch {
	case err != nil:
		return nil, err
	case len(rest) == 1:
		return d.At(issuer, doc, 0)
	case rest[1] == listFile:
		return d.Fetch(context.Background(), listPath(Key{Issuer: issuer, Doc: doc}.DocKey()))
	}
	v, ok := parseHexID(strings.TrimSuffix(rest[1], geulExt), 2)
	if !ok {
		return nil, fmt.Errorf("geul/resolve: %q is not a version ID", rest[1])
	}
	return d.At(issuer, doc, binary.BigEndian.Uint32(v))
}
