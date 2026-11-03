//ff:func feature=notation type=decoder control=sequence
//ff:what markerToken — the token of a named slot marker record (ABSENT, UNKNOWN, NONE); false for any other record
package notation

import "github.com/geul-org/gogeul/pkg/core"

// markerToken returns the markerTokens token of r when r is a named slot
// marker (core.IsSlotMarker, C 0~2; its WORD is its TAG, 0x0000/0x0010/0x0020).
func markerToken(r core.Record) (string, bool) {
	tok, ok := markerNames[core.TargetWord(r.Kind, r.Codec)]
	return tok, ok && core.IsSlotMarker(r)
}
