//ff:func feature=core type=decoder control=selection
//ff:what DecodeTargetPointer — the Issuer, document or codec ID and version ID of a DOC-POINTER or CODEC-POINTER
package core

// DecodeTargetPointer splits a DOC-POINTER (4-word document ID) or a
// CODEC-POINTER (2-word codec ID) into its Issuer, ID and 2-word version ID.
// The places that name a published document — DOC-ID, the version list
// records, BIND Subjects, IMPORT-DOC — take either; the ID width tells them
// apart.
func DecodeTargetPointer(r Record) (issuer, id, version []Word, err error) {
	switch {
	case r.Kind == KindReference && r.Codec == CodecCodecPointer:
		iss, _, _, err := DecodeCodecPointer(r)
		if err != nil {
			return nil, nil, nil, err
		}
		w := len(iss)
		return iss, r.Payload[w : w+2], r.Payload[w+2:], nil
	default:
		return DecodeDocPointer(r)
	}
}
