//ff:func feature=schema type=rule control=sequence
//ff:what globeCodec — whether a Reference C may name COORDINATE-GEO's globe: a 5Q Identified code of any scheme but WIKIDATA-PID
package schema

import "github.com/geul-org/gogeul/pkg/core"

// globeCodec reports whether c is in COORDINATE-GEO's closed globe list
// (grammar 0x065L, globe CODEC 0050 00A0 00F0 0140 0200 0250 02A0 02F0 0340
// 0390): the Identified code of every 5Q scheme except WIKIDATA-PID (a
// property is no body). Pointers, REFERENCE-ISSUER, the empty codes and the
// other Quantifiers are not.
func globeCodec(c uint16) bool {
	q, ok := core.RefQuantifier(c)
	s, _ := core.RefScheme(c)
	return ok && q == core.QuantIdentified && s != core.SchemeWikidataPID
}
