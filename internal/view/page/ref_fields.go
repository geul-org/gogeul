//ff:func feature=page type=helper control=selection
//ff:what refFields — the named fields of a Reference payload by its C: IMPORT ID, Issuer, document or codec ID, version ID, target ID
package page

import "github.com/geul-org/gogeul/pkg/core"

// refFields splits a Reference payload of n words by its C:
// EXTERNAL-POINTER-* (scheme 6~8) is IMPORT ID and target ID,
// EXTERNAL-ISSUER-* (scheme 9~11) is Issuer, document ID and target ID,
// ISSUER-POINTER is the Issuer, DOC-POINTER is Issuer, document ID and version
// ID, CODEC-POINTER is Issuer, codec ID and version ID, a REFERENCE-ISSUER
// payload is the Issuer's (one field), and the other schemes are the target ID
// only. A field cut short by a short payload keeps what is there.
func refFields(c uint16, n int) []refField {
	scheme, _ := core.RefScheme(c)
	switch {
	case core.ExternalPointerWidth(scheme) > 0:
		w := min(core.ExternalPointerWidth(scheme), n)
		return []refField{{"IMPORT ID", w}, {"대상 ID", n - w}}
	case core.ExternalIssuerWidth(scheme) > 0:
		w := min(core.ExternalIssuerWidth(scheme), n)
		d := min(4, n-w)
		return []refField{{"Issuer", w}, {"문서 ID", d}, {"대상 ID", n - w - d}}
	case c == core.CodecIssuerPointer:
		return []refField{{"Issuer", n}}
	case c == core.CodecDocPointer:
		return pointerFields(n, "문서 ID", 4)
	case c == core.CodecCodecPointer:
		return pointerFields(n, "코덱 ID", 2)
	case core.IssuerWidth(core.KindReference, c) > 0:
		return []refField{{"Payload", n}}
	default:
		return []refField{{"대상 ID", n}}
	}
}
