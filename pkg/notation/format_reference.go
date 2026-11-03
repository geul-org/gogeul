//ff:func feature=notation type=encoder control=selection
//ff:what formatReference — LOCAL, Q, P, DOC-POINTER, ISSUER-POINTER and EXTERNAL-ISSUER-* (Identified) as short tokens
package notation

import "github.com/geul-org/gogeul/pkg/core"

// formatReference writes LOCAL, the Wikidata IDs, DOC-POINTER, ISSUER-POINTER
// and EXTERNAL-ISSUER-* with Quantifier Identified. Others (slot markers,
// another Quantifier, CODEC-POINTER, REFERENCE-ISSUER) are false.
func formatReference(r core.Record) (string, bool) {
	scheme, _ := core.RefScheme(r.Codec)
	switch c := r.Codec; {
	case c == core.CodecLocal || c == core.CodecWikidataQID || c == core.CodecWikidataPID:
		return formatNumbered(c, r.Payload)
	case c == core.CodecDocPointer:
		issuer, doc, ver, err := core.DecodeDocPointer(r)
		if err != nil {
			return "", false
		}
		return docAddress(issuer, doc, uint32(ver[0])<<16|uint32(ver[1])), true
	case c == core.CodecIssuerPointer:
		issuer, err := core.DecodeIssuerPointer(r)
		if err != nil {
			return "", false
		}
		return "ISSUER-POINTER(" + core.FormatID(issuer) + ")", true
	case core.RefIdentified(c) == c && scheme >= core.SchemeExternalIssuer2Word:
		return formatExternalIssuer(r)
	default:
		return "", false
	}
}
