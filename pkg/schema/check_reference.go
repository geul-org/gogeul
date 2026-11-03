//ff:func feature=schema type=rule control=selection
//ff:what checkReference — Reference rules by C: slot markers, pointers, REFERENCE-ISSUER, EXTERNAL-ISSUER/POINTER schemes, direct targets
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkReference applies the Reference rules by C: a slot marker standing as a
// record (checkSlotMarker), the ISSUER-, DOC- and CODEC-POINTER layouts (R5),
// the REFERENCE-ISSUER codes (their Payload is the Issuer's; F5 checked the
// extension), and by scheme C/5 the EXTERNAL-ISSUER (R4), EXTERNAL-POINTER
// (R3) and direct target (R2) forms. Every Reference C is assigned.
func checkReference(r rec, imports map[uint64]bool) []Fact {
	c := r.R.Codec
	scheme, quantified := core.RefScheme(c)
	switch {
	case core.IsSlotMarkerCodec(c):
		return checkSlotMarker(r)
	case c == core.CodecIssuerPointer || c == core.CodecDocPointer || c == core.CodecCodecPointer:
		return checkIssuerDocPointer(r, c)
	case !quantified:
		return nil
	case scheme >= core.SchemeExternalIssuer2Word:
		return checkIssuerRef(r, scheme)
	case scheme >= core.SchemeExternalPointer1Word:
		return checkPointer(r, scheme, imports)
	default:
		return checkTarget(r, r.R.Payload)
	}
}
