//ff:func feature=core type=helper control=sequence
//ff:what IssuerFormatName — the spec name of a Value ISSUER format (2-WORD Issuer and codec number); "" when unknown
package core

// IssuerFormatName returns the spec name of the ISSUER format of a Value
// record r (grammar Value ISSUER 형식): its 2-WORD Issuer and codec number
// looked up in valueIssuerFormatNames. Other kinds, wider Issuers and
// unassigned formats return "".
func IssuerFormatName(r Record) string {
	if r.Kind != KindValue || len(r.Issuer) != 2 {
		return ""
	}
	return valueIssuerFormatNames[uint64(r.Issuer[0])<<48|uint64(r.Issuer[1])<<32|uint64(r.Format)]
}
