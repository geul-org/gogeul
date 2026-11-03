//ff:func feature=pack type=helper control=sequence
//ff:what parsePPath — `<문서 ID>/versions.geul` as slot 0 and `<문서 ID>/<V>.geul` (V ≠ 0) as slot V under an Issuer's docs/
package pack

import (
	"encoding/binary"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// parsePPath reads `<문서 ID>/<name>` under ref's docs/: a 16-digit document
// ID with versions.geul (slot 0) or a non-zero 8-digit version (slot V). Any
// other spelling, latest.geul included, and the reserved Issuer list ID
// (core.DocIDIssuers; the list lives in `issuers/`) leave ref outside the
// forms.
func parsePPath(ref pathRef, doc, name string) pathRef {
	d, ok := parseHex(doc, 16)
	if !ok || !storedTail(name) || binary.BigEndian.Uint64(d) == core.DocIDIssuers {
		return pathRef{}
	}
	ref.space, ref.doc = spaceP, binary.BigEndian.Uint64(d)
	if name == listFile {
		return ref
	}
	v, _ := parseHex(strings.TrimSuffix(name, ".geul"), 8)
	ref.slot = binary.BigEndian.Uint32(v)
	return ref
}
