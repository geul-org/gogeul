//ff:func feature=schema type=rule control=selection
//ff:what checkReference — Reference rules by C: slot markers, pointers, REFERENCE-ISSUER, empty codes (not reported), EXTERNAL-ISSUER/POINTER schemes, direct targets
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkReference applies the Reference rules by C: a slot marker standing as a
// record (checkSlotMarker), the ISSUER-, DOC- and CODEC-POINTER layouts (R5),
// the REFERENCE-ISSUER codes (their Payload is the Issuer's; F5 checked the
// extension), an empty code (C 0x03·0x04, 0x1E·0x1F, 0x3E·0x3F, 0x43~0x5C:
// not reported, the reserved-code Warn was removed)
// and by scheme (core.RefScheme) the EXTERNAL-ISSUER (R4),
// EXTERNAL-POINTER (R3) and direct target (R2) forms.
func checkReference(r rec, imports map[uint64]bool) []Fact {
	c := r.R.Codec
	scheme, quantified := core.RefScheme(c)
	switch {
	case core.IsSlotMarkerCodec(c):
		return checkSlotMarker(r)
	case c == core.CodecIssuerPointer || c == core.CodecDocPointer || c == core.CodecCodecPointer:
		return checkIssuerDocPointer(r, c)
	case core.IssuerWidth(core.KindReference, c) > 0:
		return nil
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
