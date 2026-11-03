//ff:func feature=notation type=decoder control=sequence
//ff:what formatExternalIssuer — an EXTERNAL-ISSUER-* Reference as <Issuer>/<문서 ID>#n
package notation

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// formatExternalIssuer writes the Issuer, the 4-word document ID and the
// LOCAL target ID. A payload too short for its Issuer width, or a target that
// does not read as #n, is false.
func formatExternalIssuer(r core.Record) (string, bool) {
	s, _ := core.RefScheme(r.Codec)
	w := core.ExternalIssuerWidth(s)
	if len(r.Payload) <= w+4 {
		return "", false
	}
	local, ok := formatNumbered(core.CodecLocal, r.Payload[w+4:])
	return docAddress(r.Payload[:w], r.Payload[w:w+4], 0) + local, ok
}
