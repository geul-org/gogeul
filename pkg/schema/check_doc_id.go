//ff:func feature=schema type=rule control=sequence
//ff:what checkDocID — M2: DOC-ID is one EXTERNAL-ISSUER-*(Issuer, document ID, 0) Reference
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkDocID applies M2 to the payload of a DOC-ID: a DOC-POINTER, or for a
// codec document a CODEC-POINTER, with version 0.
func checkDocID(r rec) []Fact {
	_, _, err := core.DecodeDocID(r.R)
	if _, _, cerr := core.DecodeDocIDCodec(r.R); err != nil && cerr != nil {
		return []Fact{factAt("M2", Error, r, "Payload", "DOC-POINTER(Issuer, 문서 ID, 0) 또는 CODEC-POINTER(Issuer, 코덱 ID, 0) Reference 하나", err.Error())}
	}
	return nil
}
