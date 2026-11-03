//ff:func feature=core type=helper control=sequence
//ff:what IsAbsent — whether a record is the ABSENT slot marker (WORD 0x0000)
package core

// IsAbsent reports whether r is the Absent record: Reference C 0, no payload
// and no ISSUER extension (WORD 0x0000).
func IsAbsent(r Record) bool {
	return r.Kind == KindReference && r.Codec == 0 && len(r.Payload) == 0 && len(r.Issuer) == 0
}
