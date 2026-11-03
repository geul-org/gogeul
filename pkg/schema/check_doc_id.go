//ff:func feature=schema type=rule control=sequence
//ff:what checkDocID — M2: DOC-ID is one DOC-POINTER(Issuer, document ID, 0) or CODEC-POINTER(Issuer, codec ID, 0) Reference, and not the reserved list ID F…F
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkDocID applies M2 to the payload of a DOC-ID: a DOC-POINTER, or for a
// codec document a CODEC-POINTER, with version 0. Its document ID is never
// 0xFFFFFFFFFFFFFFFF: that reserved ID names the Issuer list, and list
// documents have no DOC-ID (grammar 0x1C7L, 0x041L).
func checkDocID(r rec) []Fact {
	_, doc, err := core.DecodeDocID(r.R)
	if _, _, cerr := core.DecodeDocIDCodec(r.R); err != nil && cerr != nil {
		return []Fact{factAt("M2", Error, r, "Payload", "DOC-POINTER(Issuer, 문서 ID, 0) 또는 CODEC-POINTER(Issuer, 코덱 ID, 0) Reference 하나", err.Error())}
	}
	if err == nil && core.FormatID(doc) == reservedListID {
		return []Fact{factAt("M2", Error, r, "문서 ID", "예약 문서 ID 0xFFFFFFFFFFFFFFFF가 아님(목록 문서는 DOC-ID가 없다)", reservedListID)}
	}
	return nil
}
