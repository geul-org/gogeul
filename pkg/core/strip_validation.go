//ff:func feature=core type=helper control=iteration dimension=1
//ff:what stripValidation — copy a record run without the validation-class Meta records (Meta C >= 0x25)
package core

// stripValidation copies the records of ws, a run already checked by
// SplitDocument, leaving out every validation-class Meta record: Kind Meta and
// C >= 0x25 (IsValidationMeta), the codes ISSUER-KEY onward and the empty
// codes after them.
func stripValidation(ws []Word) []Word {
	out := make([]Word, 0, len(ws))
	for i := 0; i < len(ws); {
		r, m, _ := DecodeRecord(ws[i:])
		if !IsValidationMeta(r.Kind, r.Codec) {
			out = append(out, ws[i:i+m]...)
		}
		i += m
	}
	return out
}
